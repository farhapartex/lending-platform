# Backend API Reference

Base path: **`/api/v1`** · JSON only · Local: `http://localhost:8080`

## Scope

This backend exists to serve the two things a blockchain is bad at: **history** and **cross-user
aggregation**. Everything else the UI needs it reads straight from the contracts.

So the surface is deliberately small — **six endpoints plus a health check.**

| Endpoint | Method | Serves | Status |
| --- | --- | --- | --- |
| [`/health`](#get-health) | GET | Container healthcheck, uptime monitoring | **Live** |
| [`/accounts/{address}/transactions`](#get-accountsaddresstransactions) | GET | `/history` table, filters, pagination | Planned |
| [`/accounts/{address}/transactions/{transactionId}`](#get-accountsaddresstransactionstransactionid) | GET | `TxDetailDrawer` | Planned |
| [`/accounts/{address}/activity`](#get-accountsaddressactivity) | GET | Dashboard `RecentActivityList` | Planned |
| [`/liquidations/eligible`](#get-liquidationseligible) | GET | `LiquidatablePositionsTable` | Planned |
| [`/liquidations/history`](#get-liquidationshistory) | GET | Past liquidations | Planned |
| [`/liquidations/{liquidationId}`](#get-liquidationsliquidationid) | GET | `LiquidationReceipt` | Planned |

**Only `/health` is implemented.** The six planned endpoints all serve indexed event history, so they are
blocked on the same thing: **the indexer**. None can be built before it exists.

Every endpoint is a **GET**. Nothing takes a request body — input travels in the path or query string.
Request payloads will be documented here if a POST is ever added.

---

## Conventions

### Authentication

**There is none in Phase 1.** The wallet address is a plain path parameter and is trusted.

That is safe here because all six endpoints serve **data already public on the blockchain** — anyone can read
any address's history straight from the chain. Guarding it would add friction and no privacy.

It would *not* be safe for anything storing off-chain personal data (an alert email, notification
preferences). Those endpoints are deliberately absent until Phase 2 adds signature-based auth.

### Headers

| Header | Direction | Notes |
| --- | --- | --- |
| `X-Request-Id` | request (optional) | Echoed back. Generated if you do not send one |
| `X-Request-Id` | response | Always present. Quote it when reporting a problem |
| `Content-Type` | response | `application/json; charset=utf-8` |

### Amounts are strings, never JSON numbers

A token amount is a `uint256`. It does not fit in a JSON number, and JavaScript would silently round it.
So every amount is an object:

```json
{ "amount": "50132158115", "decimals": 6, "symbol": "USDC" }
```

`amount` is the raw base-unit value **as a string**. Divide by `10^decimals` to display —
`50132158115 / 10^6` = `50,132.158115 USDC`.

The same applies to USD values, which use 8 decimals: `"580000000000"` is `$5,800.00`.

### IDs are masked

Public IDs are never the database's sequential primary key — that leaks row counts and invites enumeration.
Each is an opaque, prefixed token:

```
txn_mvtep746x22iqzn52kca     a transaction
liq_...                      a liquidation
```

The prefix identifies the resource type, so passing one where another is expected is rejected rather than
silently reading the wrong row. Treat them as opaque — never parse or construct one.

### Error envelope

Every error, on every endpoint:

```json
{
  "error": { "code": "NOT_FOUND", "message": "The requested endpoint does not exist." },
  "request_id": "9949a5499f6ed44992f321a3"
}
```

`details` is an optional object of field-level context, omitted when empty.

**Branch on `code`, never on `message`.** Messages are written for humans and may be reworded; codes are a
stable set.

| Code | HTTP | Meaning |
| --- | --- | --- |
| `BAD_REQUEST` | 400, 405 | Input or method was not usable |
| `NOT_FOUND` | 404 | Endpoint or resource does not exist |
| `INTERNAL_ERROR` | 500, 503 | Something failed on our side |

More codes will be added as endpoints land, for example `INVALID_ADDRESS`.

### Lists and pagination

Cursor-based, not page numbers — an offset makes the database walk every skipped row, while a cursor is an
index seek however deep you go.

```json
{
  "items": [],
  "next_cursor": "b3BhcXVl",
  "as_of": { "block": 1234, "time": "2026-08-19T03:25:46Z" }
}
```

- `next_cursor` is opaque. Pass it back as `?cursor=`. `null` means the end
- `as_of` says how fresh the data is, so the UI can show "as of 3 blocks ago" rather than implying live
- `limit` defaults to 25 and is capped server-side

---

# Live

## GET /health

Liveness and build information. Used by the Docker healthcheck and uptime monitoring. Takes no parameters.

```bash
curl -s http://localhost:8080/api/v1/health
```

### Response — 200 OK

```json
{
  "status": "ok",
  "service": "core-service",
  "version": "dev",
  "environment": "local",
  "uptime_seconds": 61,
  "checked_at": "2026-08-19T03:25:46Z"
}
```

| Field | Type | Notes |
| --- | --- | --- |
| `status` | string | `ok`, `degraded`, or `down` |
| `service` | string | From `SERVICE_NAME` |
| `version` | string | Build version, injected at compile time |
| `environment` | string | `local`, `staging`, or `production` |
| `uptime_seconds` | number | Whole seconds since process start |
| `checked_at` | string | RFC 3339, UTC |

### Response — 503 Service Unavailable

Returned when the health check itself cannot complete.

```json
{
  "error": { "code": "INTERNAL_ERROR", "message": "Health check could not be completed." },
  "request_id": "..."
}
```

**Note:** this reports that the *process* is alive. It does not check Postgres, Redis, or the chain node. A
dependency check is worth adding, but a health endpoint that fails when a cache is down can make an
orchestrator restart a service that was serving traffic perfectly well — so it needs splitting into liveness
and readiness rather than bolting checks onto this one.

---

# Planned

All six depend on the indexer. Shapes below are the agreed contract.

## GET /accounts/{address}/transactions

The `/history` table. Cached ~5s.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `address` | path | — | `0x` + 40 hex, case-insensitive. Send lowercase to keep cache keys stable |
| `kind` | query | all | `deposit`, `withdraw`, `borrow`, `repay`, `collateral_added`, `collateral_withdrawn`, `liquidation` |
| `from`, `to` | query | — | RFC 3339 bounds. Send UTC — a date picker's local midnight is not UTC midnight |
| `cursor` | query | — | From a previous `next_cursor` |
| `limit` | query | 25 | Capped at 100 |

### Response — 200 OK

```json
{
  "items": [
    {
      "id": "txn_mvtep746x22iqzn52kca",
      "kind": "borrow",
      "amount": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
      "tx_hash": "0x9f2c…",
      "block": 42,
      "block_time": "2026-08-19T03:20:11Z",
      "status": "confirmed"
    }
  ],
  "next_cursor": "b3BhcXVl",
  "as_of": { "block": 45, "time": "2026-08-19T03:25:46Z" }
}
```

Ordered newest first.

**An unknown address returns 200 with an empty list, not 404.** Every address is valid on a blockchain
whether or not it has ever transacted, so "no history" is an empty result, not a missing resource.

**There is a lag between a user signing and this endpoint knowing.** The indexer waits for confirmations and
then polls, so a transaction is visible on-chain for roughly 10–30 seconds before it appears here. The UI
bridges that by tracking its own pending transactions locally and dropping each one when its `tx_hash`
appears in a page from this endpoint.

Errors: `400 BAD_REQUEST` for a malformed address or cursor.

---

## GET /accounts/{address}/transactions/{transactionId}

One transaction in full, for the detail drawer. Cached ~30s.

Adds to the list item: the resulting health factor, gas used, and — for a liquidation — the collateral
seized and bonus paid.

Errors: `404 NOT_FOUND` if the transaction does not exist **or does not belong to that address**.

---

## GET /accounts/{address}/activity

The dashboard's short recent-activity list. Same item shape as `/transactions`, no cursor. Cached ~5s.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `limit` | query | 5 | Capped at 20 |

---

## GET /liquidations/eligible

Positions that can be liquidated right now. **Deliberately public and unauthenticated** — liquidators must
be able to find work before connecting a wallet. Cached ~5s, rate limited by IP.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `sort` | query | `health` | `health`, `size`, `reward` |
| `cursor`, `limit` | query | 25 | Capped at 100 |

### Response — 200 OK

```json
{
  "items": [
    {
      "borrower": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
      "health_factor_bps": 9098,
      "debt_to_repay": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
      "collateral_to_seize": { "amount": "1846551724137931035", "decimals": 18, "symbol": "WETH" },
      "bonus_value_usd": "25500000000",
      "shortfall_value_usd": "0"
    }
  ],
  "next_cursor": null,
  "as_of": { "block": 45, "time": "2026-08-19T03:25:46Z" }
}
```

Two things a liquidation bot must understand:

**`shortfall_value_usd` above zero means the liquidation loses money.** The collateral is worth less than the
debt plus bonus, so you would pay more than you receive. Filter these out unless you have a reason not to.

**This list can be stale.** Always call `previewLiquidation` on the contract before submitting, and expect
`PositionIsHealthy` sometimes — it means another liquidator got there first, or the price recovered.

---

## GET /liquidations/history

Completed liquidations, public. Cached ~15s. Same pagination as `eligible`. Each item adds `liquidator`,
`tx_hash`, `block_time`, and the health factor before the liquidation.

---

## GET /liquidations/{liquidationId}

One liquidation receipt: debt repaid, collateral seized, bonus, health factor before, the trigger price and
its decimals, and any shortfall. Cached ~60s.

Every field comes from the `LiquidationExecuted` event, which was designed to carry all of it — so building
a receipt never needs a follow-up call to the chain.

---

## Endpoints we are deliberately not building

Listed so nobody adds them back thinking they were forgotten.

### Because a contract call already answers it

| Endpoint | Replaced by |
| --- | --- |
| `/protocol/stats` | `PositionLens.marketData()` — same totals, one call |
| `/markets`, `/markets/{id}` | `PositionLens.marketData()` |
| `/accounts/{address}/summary` | `PositionLens.accountData(address)` |
| `/accounts/{address}/positions` | `PositionLens.accountData(address)` |

The last two are not merely redundant, they are **hazardous**. They would return health factor, borrow limit
and withdrawable collateral — figures the UI uses to size a transaction the user is about to sign. Served
from the backend they are blocks out of date, so a borrow sized from them reverts. The rule is *the contracts
decide, everything else displays*, and these endpoints invite breaking it.

### Because nothing consumes it yet

| Endpoint | Reason |
| --- | --- |
| `/markets/{id}/snapshots` | Rate-history charts. No chart component exists in the frontend |

### Because they need Phase 2 auth

| Endpoint | Reason |
| --- | --- |
| `/accounts/{address}/notifications` | Read state is off-chain |
| `/accounts/{address}/preferences` | Stores an email address |
| `/accounts/{address}/email`, `/email/verify` | Same |
| `/faucet/status`, `/faucet/requests` | A rate-limited resource someone could drain on your behalf |
| `/stream` (SSE) | Live push. Polling is sufficient for now |

Building any of these before auth exists would ship a window in which anyone can write to anyone's record.

---

## Trying it

```bash
curl -s http://localhost:8080/api/v1/health | python3 -m json.tool

curl -s -H "X-Request-Id: my-trace-id" -D - -o /dev/null \
  http://localhost:8080/api/v1/health

curl -s http://localhost:8080/api/v1/nope
```
