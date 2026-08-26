# Backend API Reference

Base path: **`/api/v1`** · JSON only · Local: `http://localhost:8080`

## Scope

This backend exists to serve the two things a blockchain is bad at: **history** and **cross-user
aggregation**. Everything else the UI needs it reads straight from the contracts.

So the surface is deliberately small — **six endpoints plus a health check.**

| Endpoint | Method | Serves | Status |
| --- | --- | --- | --- |
| [`/health`](#get-health) | GET | Container healthcheck, uptime monitoring | **Live** |
| [`/accounts/{address}/transactions`](#get-accountsaddresstransactions) | GET | `/history` table, filters, pagination | **Live** |
| [`/accounts/{address}/transactions/{transactionId}`](#get-accountsaddresstransactionstransactionid) | GET | `TxDetailDrawer` | **Live** |
| [`/accounts/{address}/activity`](#get-accountsaddressactivity) | GET | Dashboard `RecentActivityList` | **Live** |
| [`/liquidations/eligible`](#get-liquidationseligible) | GET | `LiquidatablePositionsTable` | Planned |
| [`/liquidations/history`](#get-liquidationshistory) | GET | Past liquidations | **Live** |
| [`/liquidations/{liquidationId}`](#get-liquidationsliquidationid) | GET | `LiquidationReceipt` | **Live** |

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

### Browser origins

The API is called directly from the browser, so it carries an origin allowlist. `CORS_ALLOWED_ORIGINS` is a
comma-separated list and defaults to `http://localhost:5173` locally and to nothing at all elsewhere — a
deployed environment must name its own origin. A `*` is accepted locally and refused in production.

An allowed origin gets `Access-Control-Allow-Origin` on **every** response including errors, which matters:
without it the browser blocks the body and a clean `404` looks like a network failure. `X-Request-Id` is
listed in `Access-Control-Expose-Headers` so client code can read it. Credentials are never allowed —
Phase 1 sends no cookies and no auth headers.

An origin that is not on the list still gets served, just without the CORS headers, so the browser is the
thing that blocks it. Preflights from an unknown origin are refused with `403`.

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

## GET /accounts/{address}/transactions

The `/history` table. Ordered newest first.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `address` | path | — | `0x` + 40 hex, case-insensitive. Send lowercase to keep client cache keys stable |
| `kind` | query | all | One or more of `deposit`, `withdraw`, `borrow`, `repay`, `collateral_added`, `collateral_withdrawn`, `liquidation`. Repeat the parameter or comma-separate; duplicates are ignored |
| `from`, `to` | query | — | RFC 3339 bounds, inclusive. Send UTC — a date picker's local midnight is not UTC midnight |
| `cursor` | query | — | An opaque `next_cursor` from a previous response. Never build one |
| `limit` | query | 25 | Capped at 100. `0` or absent means the default |

```bash
ADDRESS=0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
BASE=http://localhost:8080/api/v1/accounts/$ADDRESS/transactions

curl -s "$BASE?limit=3"
curl -s "$BASE?kind=borrow,repay"
curl -s "$BASE?from=2026-08-22T07:00:00Z&to=2026-08-22T09:00:00Z"
```

### Response — 200 OK

```json
{
  "items": [
    {
      "id": "txn_mvtep746x22iqzn52kca",
      "kind": "borrow",
      "amount": { "amount": "2000000", "decimals": 6, "symbol": "USDC" },
      "health_factor_after_bps": 12002,
      "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000002",
      "block": 42,
      "block_time": "2026-08-22T07:00:00Z",
      "log_index": 2,
      "status": "confirmed"
    }
  ],
  "next_cursor": "djEuMTc4NzM3ODQwMDAwMDAwMC4z",
  "as_of": { "block": 4218, "time": "2026-08-22T09:30:00Z" }
}
```

Each item has exactly the same shape as the detail endpoint below, so one client type covers both.

| Field | Type | Notes |
| --- | --- | --- |
| `items` | array | Always an array, never `null`, so a client can render without a null check |
| `next_cursor` | string or null | Pass it back as `cursor` for the next page. **Null means there is genuinely nothing more** |
| `as_of.block` | number or null | The last block the indexer has processed. Null until the indexer has run at all — it is never invented from the rows returned |
| `as_of.time` | string | When the indexer last advanced, not when you asked. A timestamp drifting into the past is how you spot a stalled indexer |

### Paging

Keyset paging on `(block_time, id)`, not offsets — a new transaction arriving mid-scroll cannot shift rows
onto a page you already read.

**A null `next_cursor` is trustworthy.** The query fetches one row beyond the page and drops it, so a
cursor is only ever returned when a further row really exists. A final page that happens to be exactly
`limit` rows long returns `null`, not a cursor leading to an empty page.

A cursor carries only a position, never an identity, and every query is scoped to the wallet in the path.
Replaying one wallet's cursor against another wallet's URL returns that wallet's own rows — it cannot leak
anything.

### Errors

| Status | Code | Cause |
| --- | --- | --- |
| 400 | `BAD_REQUEST` | A malformed address, an unknown `kind`, a `from`/`to` that is not RFC 3339, a negative or non-numeric `limit`, or a `cursor` that did not come from us. The message names the offending parameter |
| 400 | `BAD_REQUEST` | `from` is later than `to` |

**An unknown address returns 200 with an empty list, not 404.** Every address is valid on a blockchain
whether or not it has ever transacted, so "no history" is an empty result rather than a missing resource.
The `as_of` block is still reported, so the client can tell "nothing happened" from "nothing indexed yet".

A `limit` above 100 is silently capped rather than refused — asking for too much is not a client error.

**There is a lag between a user signing and this endpoint knowing.** The indexer waits for confirmations
and then polls, so a transaction is visible on-chain for roughly 10–30 seconds before it appears here. The
UI bridges that by tracking its own pending transactions locally and dropping each one when its `tx_hash`
appears in a page from this endpoint.

**This returns an empty list for every wallet until the indexer ships.** `user_transactions` has no rows
yet.

---

## GET /accounts/{address}/transactions/{transactionId}

One transaction in full, for the detail drawer.

| Parameter | In | Notes |
| --- | --- | --- |
| `address` | path | `0x` + 40 hex, case-insensitive. Checksummed and lowercase both resolve to the same wallet |
| `transactionId` | path | A masked `txn_…` id, as returned by the list endpoint |

```bash
curl -s http://localhost:8080/api/v1/accounts/0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266/transactions/txn_mvtep746x22iqzn52kca
```

### Response — 200 OK

```json
{
  "id": "txn_mvtep746x22iqzn52kca",
  "kind": "borrow",
  "amount": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
  "health_factor_after_bps": 12661,
  "tx_hash": "0x9f2ccafe",
  "block": 42,
  "block_time": "2026-08-20T03:20:11Z",
  "log_index": 3,
  "status": "confirmed"
}
```

| Field | Type | Notes |
| --- | --- | --- |
| `id` | string | The same masked id that was requested |
| `kind` | string | `deposit`, `withdraw`, `borrow`, `repay`, `collateral_added`, `collateral_withdrawn`, or `liquidation` |
| `amount` | object | Base units as a string, plus the asset's `decimals` and `symbol` for formatting |
| `health_factor_after_bps` | number or null | The borrower's health factor right after this transaction, in basis points. `12661` means 1.2661. Null when the indexer had no health reading for the event |
| `tx_hash` | string | The on-chain transaction hash |
| `block` | number | Block number the event was mined in |
| `block_time` | string | RFC 3339, UTC |
| `log_index` | number | Position of the event within the transaction. A single transaction can emit several, so `tx_hash` alone does not identify a row |
| `status` | string | Always `confirmed`. Only indexed events are stored, and the indexer waits for confirmations before writing, so nothing pending or failed can reach this endpoint |

`amount` carries `decimals` and `symbol` so the client never has to look up asset metadata to render a
figure. If the asset row is missing, `decimals` is `0` and `symbol` is empty rather than the response
failing — a detail drawer with an unformatted number beats a broken drawer.

### Errors

| Status | Code | Cause |
| --- | --- | --- |
| 400 | `BAD_REQUEST` | The id is not a well-formed `txn_…` token — wrong shape, bad tag, or the id of another kind such as `mkt_…`. Rejected before any database work |
| 400 | `BAD_REQUEST` | The address is not a valid `0x` + 40 hex string |
| 404 | `NOT_FOUND` | No such transaction, the address has no history at all, or **the transaction belongs to a different wallet** |

**Another wallet's transaction returns 404, not 403.** Guessing a valid id should not confirm that the id
exists, so a row that is not yours is indistinguishable from a row that does not exist.

**This returns 404 for every real request until the indexer ships.** `user_transactions` has no rows yet.
The endpoint, the ownership check and the masking are all live and verified against Postgres — there is
simply nothing to find.

---

## GET /accounts/{address}/activity

The dashboard's short recent-activity list. Newest first, no filters, no paging.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `address` | path | — | `0x` + 40 hex, case-insensitive |
| `limit` | query | 5 | Capped at 20. `0` or absent means the default |

```bash
ADDRESS=0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

curl -s http://localhost:8080/api/v1/accounts/$ADDRESS/activity
curl -s "http://localhost:8080/api/v1/accounts/$ADDRESS/activity?limit=10"
```

### Response — 200 OK

```json
{
  "items": [
    {
      "id": "txn_buckzju36csnq4qecunq",
      "kind": "borrow",
      "amount": { "amount": "1000000", "decimals": 6, "symbol": "USDC" },
      "health_factor_after_bps": 12001,
      "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000001",
      "block": 41,
      "block_time": "2026-08-22T08:00:00Z",
      "log_index": 1,
      "status": "confirmed"
    }
  ],
  "as_of": { "block": 4218, "time": "2026-08-22T09:30:00Z" }
}
```

**Items are byte-for-byte the same shape as `/transactions`**, so one client type covers both — they are
built by the same code path, and a test pins them together.

**There is no `next_cursor` field at all**, not even a null one. This endpoint does not page: a dashboard
panel that showed a "load more" button would be a different feature. Use `/transactions` for anything
scrollable.

`as_of` behaves exactly as on `/transactions`: the block is the indexer's last processed block, `null` until
the indexer has run, and never inferred from the rows returned.

### Errors

| Status | Code | Cause |
| --- | --- | --- |
| 400 | `BAD_REQUEST` | A malformed address, or a `limit` that is not a non-negative whole number |

**An unknown address returns 200 with an empty list**, with `as_of` still populated — so the dashboard can
tell "this wallet has done nothing" from "nothing has been indexed yet". Those need different empty states.

`kind`, `from`, `to` and `cursor` are **ignored rather than refused** here. They are not part of this
endpoint, and rejecting an unknown query parameter would make a harmless client mistake look like a failure.

A `limit` above 20 is silently capped.

---

## GET /liquidations/history

Completed liquidations, newest first. **Public and unauthenticated** — a liquidation is a public on-chain
event, and liquidators need to study past work before connecting a wallet.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `market` | query | all | A masked `mkt_…` id. Never construct one |
| `cursor` | query | — | An opaque `next_cursor` from a previous response |
| `limit` | query | 25 | Capped at 100 |

```bash
BASE=http://localhost:8080/api/v1/liquidations

curl -s "$BASE/history?limit=2"
curl -s "$BASE/history?market=mkt_zbgqajqihcctneocuahq"
```

### Response — 200 OK

```json
{
  "items": [
    {
      "id": "liq_5h74y4qafdaywxk7z64a",
      "borrower": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
      "liquidator": "0x70997970c51812dc3a010c7d01b50e0d17dc79c8",
      "debt_repaid": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
      "collateral_seized": { "amount": "1846551724137931035", "decimals": 18, "symbol": "WETH" },
      "bonus_value": { "amount": "25500000000", "decimals": 8, "symbol": "USD" },
      "shortfall_value": { "amount": "0", "decimals": 8, "symbol": "USD" },
      "health_factor_before_bps": 9099,
      "trigger_price": { "amount": "220000000000", "decimals": 8, "symbol": "USD" },
      "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000001",
      "block": 4001,
      "block_time": "2026-08-22T08:00:00Z"
    }
  ],
  "next_cursor": "djEuMTc4NzM4MjAwMDAwMDAwMC4yNg",
  "as_of": { "block": 4218, "time": "2026-08-22T09:30:00Z" }
}
```

### Three amounts, three different scales

This is the part to get right. Every figure carries its own `decimals` and `symbol`, so **never assume a
scale** — read it from the field.

| Field | Denominated in | Why |
| --- | --- | --- |
| `debt_repaid` | The market's **debt** asset | The liquidator repaid the borrower's loan, so this is USDC at 6 decimals |
| `collateral_seized` | The market's **collateral** asset | They received collateral, so this is WETH at 18 decimals |
| `bonus_value` | **USD value**, at the trigger price's scale | The contract emits the bonus as a value, not a token amount |
| `shortfall_value` | **USD value**, same scale | Also a value |
| `trigger_price` | **USD**, scale from the event | The collateral price that made the position liquidatable |

The three USD figures share the price scale because that is how the contract computes them — a value is
`amount × price ÷ 10^tokenDecimals`, so the value inherits the price's decimals. The scale comes from the
event itself rather than a constant on our side, which is why `decimals` is reported per field instead of
documented as "always 8".

**`shortfall_value` above zero means the liquidation lost money.** The collateral was worth less than the
debt plus bonus, so the liquidator paid more than they received. A bad-debt event, and worth surfacing.

### Other fields

| Field | Type | Notes |
| --- | --- | --- |
| `borrower`, `liquidator` | string | Lowercase addresses. Checksum them client-side for display |
| `health_factor_before_bps` | number or null | The borrower's health factor immediately before, in basis points. `9099` means 0.9099 — below 1.0, which is why it was liquidatable |
| `as_of` | object | As on `/transactions`: the indexer's last processed block, `null` until it has run |

Paging is keyset on `(block_time, id)` with the same guarantee as `/transactions` — a `null` `next_cursor`
genuinely means there is nothing more, because the query fetches one row beyond the page and drops it.

### Errors

| Status | Code | Cause |
| --- | --- | --- |
| 400 | `BAD_REQUEST` | `market` is not a well-formed `mkt_…` id — including a raw number or an id of another kind |
| 400 | `BAD_REQUEST` | A `limit` that is not a non-negative whole number, or a `cursor` that did not come from us |

A `market` id that is well-formed but matches no market returns **200 with an empty list**, not 404 — the
filter simply selected nothing.

---

## GET /liquidations/{liquidationId}

One liquidation receipt. Public.

| Parameter | In | Notes |
| --- | --- | --- |
| `liquidationId` | path | A masked `liq_…` id from the history list |

```bash
curl -s http://localhost:8080/api/v1/liquidations/liq_5h74y4qafdaywxk7z64a
```

### Response — 200 OK

**Byte-for-byte the same object as a history item** — one client type covers both. Every field comes from
the `LiquidationExecuted` event, which was designed to carry all of it, so building a receipt never needs a
follow-up call to the chain.

```json
{
  "id": "liq_5h74y4qafdaywxk7z64a",
  "borrower": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
  "liquidator": "0x70997970c51812dc3a010c7d01b50e0d17dc79c8",
  "debt_repaid": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
  "collateral_seized": { "amount": "1846551724137931035", "decimals": 18, "symbol": "WETH" },
  "bonus_value": { "amount": "25500000000", "decimals": 8, "symbol": "USD" },
  "shortfall_value": { "amount": "0", "decimals": 8, "symbol": "USD" },
  "health_factor_before_bps": 9099,
  "trigger_price": { "amount": "220000000000", "decimals": 8, "symbol": "USD" },
  "tx_hash": "0x0000000000000000000000000000000000000000000000000000000000000001",
  "block": 4001,
  "block_time": "2026-08-22T08:00:00Z"
}
```

**There is no `as_of` here.** A receipt is a finished historical record, so reporting indexer progress
alongside it would say nothing useful.

### Errors

| Status | Code | Cause |
| --- | --- | --- |
| 400 | `BAD_REQUEST` | The id is not a well-formed `liq_…` token, including a raw number or a `txn_…` id |
| 404 | `NOT_FOUND` | No such liquidation |

---

# Planned

This one depends on the indexer and on live position data. The shape below is the agreed contract.

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

The transaction endpoint, with a wallet from the local Anvil node. The first id is well-formed and simply
has no row behind it yet, the second is not a valid token at all:

```bash
ADDRESS=0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

curl -s http://localhost:8080/api/v1/accounts/$ADDRESS/transactions
curl -s "http://localhost:8080/api/v1/accounts/$ADDRESS/transactions?kind=borrow&limit=5"
curl -s http://localhost:8080/api/v1/accounts/$ADDRESS/activity
curl -s http://localhost:8080/api/v1/accounts/$ADDRESS/transactions/txn_mvtep746x22iqzn52kca
curl -s http://localhost:8080/api/v1/accounts/$ADDRESS/transactions/77
```

Masked ids are derived from `ID_MASK_SECRET`, so a `txn_…` token minted against one secret is rejected
against another. The examples here assume the local default.
