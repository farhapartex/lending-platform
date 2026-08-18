# Backend API Reference

Base path: **`/api/v1`** · JSON only · Local: `http://localhost:8080`

## Status at a glance

| Endpoint | Method | Status |
| --- | --- | --- |
| [`/health`](#get-health) | GET | **Live** |
| [`/protocol/stats`](#get-protocolstats) | GET | Planned |
| [`/markets`](#get-markets) | GET | Planned |
| [`/markets/{marketId}`](#get-marketsmarketid) | GET | Planned |
| [`/markets/{marketId}/snapshots`](#get-marketsmarketidsnapshots) | GET | Planned |
| [`/accounts/{address}/summary`](#get-accountsaddresssummary) | GET | Planned |
| [`/accounts/{address}/positions`](#get-accountsaddresspositions) | GET | Planned |
| [`/accounts/{address}/transactions`](#get-accountsaddresstransactions) | GET | Planned |
| [`/accounts/{address}/transactions/{transactionId}`](#get-accountsaddresstransactionstransactionid) | GET | Planned |
| [`/accounts/{address}/activity`](#get-accountsaddressactivity) | GET | Planned |
| [`/liquidations/eligible`](#get-liquidationseligible) | GET | Planned |
| [`/liquidations/history`](#get-liquidationshistory) | GET | Planned |
| [`/liquidations/{liquidationId}`](#get-liquidationsliquidationid) | GET | Planned |

**Only `/health` is implemented today.** Everything else is the agreed contract for the work in progress —
shapes may shift slightly as each one lands, and this file gets updated when it does.

Every endpoint in Phase 1 is a **GET**, so none takes a request body. Where a request needs input it comes
from the path or the query string. Request payloads will be documented here as soon as a POST exists.

---

## Conventions

### Authentication

**There is none in Phase 1.** The wallet address is passed as a plain path parameter and trusted.

That is safe for these endpoints because every one of them serves **data already public on the
blockchain** — anyone can read any address's positions and history straight from the chain. Guarding them
would add friction and no privacy.

It is *not* safe for anything storing off-chain personal data (an alert email, notification preferences), so
those endpoints are deliberately **not built** until Phase 2 brings signature-based auth.

### Headers

| Header | Direction | Notes |
| --- | --- | --- |
| `X-Request-Id` | request (optional) | Echoed back. If you do not send one, the server generates it |
| `X-Request-Id` | response | Always present. Quote it when reporting a problem |
| `Content-Type` | response | `application/json; charset=utf-8` |

### Amounts are strings, never JSON numbers

A token amount is a `uint256`. It does not fit in a JSON number, and JavaScript would silently round it.
So every amount is an object:

```json
{
  "amount": "50132158115",
  "decimals": 6,
  "symbol": "USDC"
}
```

`amount` is the **raw base-unit value as a string**. Divide by `10^decimals` to display it — here
`50132158115 / 10^6` = `50,132.158115 USDC`.

The same rule applies to prices (8 decimals) and indexes (18 decimals).

### IDs are masked

Public IDs are never the database's sequential primary key — that would leak how many rows exist and invite
enumeration. Each ID is an opaque, prefixed token:

```
mkt_zuxjejolmpb2ynh4g32q     a market
txn_mvtep746x22iqzn52kca     a transaction
liq_...                      a liquidation
```

The prefix identifies the resource type, so passing a market ID where a transaction ID is expected is
rejected rather than silently reading the wrong row. Treat them as opaque strings — do not parse them.

### Error envelope

Every error, on every endpoint, has the same shape:

```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "The requested endpoint does not exist."
  },
  "request_id": "9949a5499f6ed44992f321a3"
}
```

`details` is an optional object of extra field-level context, omitted when empty.

**Branch on `code`, never on `message`.** Messages are written for humans and may be reworded; codes are a
stable set.

| Code | HTTP | Meaning |
| --- | --- | --- |
| `BAD_REQUEST` | 400, 405 | Input or method was not usable |
| `NOT_FOUND` | 404 | Endpoint or resource does not exist |
| `INTERNAL_ERROR` | 500, 503 | Something failed on our side |

More codes will be added as endpoints land (for example `MARKET_NOT_FOUND`, `INVALID_ADDRESS`).

### Lists and pagination

List endpoints use **cursor** pagination, not page numbers — an offset makes the database walk every skipped
row, while a cursor is a direct index seek however deep you go.

```json
{
  "items": [],
  "next_cursor": "b3BhcXVl",
  "as_of": {
    "block": 1234,
    "time": "2026-08-18T03:25:46Z"
  }
}
```

- `next_cursor` is opaque. Pass it back as `?cursor=` for the next page. `null` means you have reached the end
- `as_of` tells you how fresh the data is, so the UI can say "as of 3 blocks ago" rather than implying it is live
- `limit` defaults to 25 and is capped server-side

---

# Live endpoints

## GET /health

Liveness and build information. Used by Docker's healthcheck and by uptime monitoring.

Takes no parameters.

### Request

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
  "checked_at": "2026-08-18T03:25:46Z"
}
```

| Field | Type | Notes |
| --- | --- | --- |
| `status` | string | `ok`, `degraded`, or `down` |
| `service` | string | Service name, from `SERVICE_NAME` |
| `version` | string | Build version, injected at compile time |
| `environment` | string | `local`, `staging`, or `production` |
| `uptime_seconds` | number | Whole seconds since the process started |
| `checked_at` | string | RFC 3339, UTC |

### Response — 503 Service Unavailable

Returned when the health check itself cannot complete.

```json
{
  "error": {
    "code": "INTERNAL_ERROR",
    "message": "Health check could not be completed."
  },
  "request_id": "..."
}
```

**Note:** this endpoint reports that the *process* is alive. It does not currently check Postgres, Redis, or
the chain node. A dependency check is worth adding, but a health endpoint that fails when a cache is down
can cause an orchestrator to restart a service that was serving traffic perfectly well — so the checks need
to be split into liveness and readiness rather than bolted onto this one.

---

# Planned endpoints

Shapes below are the agreed contract. None is implemented yet.

## GET /protocol/stats

Protocol-wide totals for the landing page strip and the top of `/markets`. Cached ~15s.

### Response — 200 OK

```json
{
  "chain_id": 31337,
  "total_supplied": { "amount": "150000000000", "decimals": 6, "symbol": "USDC" },
  "total_borrowed": { "amount": "15000000000", "decimals": 6, "symbol": "USDC" },
  "available_liquidity": { "amount": "135000000000", "decimals": 6, "symbol": "USDC" },
  "utilization_bps": 1000,
  "market_count": 1,
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

`utilization_bps` is basis points — `1000` means 10.00%.

---

## GET /markets

Every market on the configured chain. Cached ~10s.

### Response — 200 OK

```json
{
  "items": [
    {
      "id": "mkt_zuxjejolmpb2ynh4g32q",
      "chain_id": 31337,
      "status": "active",
      "collateral_asset": {
        "symbol": "WETH",
        "name": "Wrapped Ether",
        "decimals": 18,
        "address": "0x5fbdb2315678afecb367f032d93f642f64180aa3"
      },
      "debt_asset": {
        "symbol": "USDC",
        "name": "USD Coin",
        "decimals": 6,
        "address": "0xe7f1725e7734ce288f8367e1bb143e90bb3f0512"
      },
      "total_supplied": { "amount": "150000000000", "decimals": 6, "symbol": "USDC" },
      "total_borrowed": { "amount": "15000000000", "decimals": 6, "symbol": "USDC" },
      "available_liquidity": { "amount": "135000000000", "decimals": 6, "symbol": "USDC" },
      "utilization_bps": 1000,
      "supply_apr_bps": 16,
      "borrow_apr_bps": 181,
      "max_ltv_bps": 7500,
      "liquidation_threshold_bps": 8000,
      "liquidation_bonus_bps": 500,
      "reserve_factor_bps": 1000,
      "kink_utilization_bps": 8000,
      "min_deposit": { "amount": "1000000", "decimals": 6, "symbol": "USDC" },
      "deposits_paused": false,
      "borrow_paused": false
    }
  ],
  "next_cursor": null,
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

`status` is one of `active`, `deposit_only`, `deprecated`.

Note `max_ltv_bps` (7500) and `liquidation_threshold_bps` (8000) are deliberately different. You may borrow
up to 75% of your collateral, but you are only liquidated at 80%. That 5-point gap is the safety buffer.

---

## GET /markets/{marketId}

One market, same object as a `/markets` list item. Cached ~10s.

| Parameter | In | Notes |
| --- | --- | --- |
| `marketId` | path | Masked ID, `mkt_...` |

Errors: `404 NOT_FOUND` for an unknown market, `400 BAD_REQUEST` for a malformed ID.

---

## GET /markets/{marketId}/snapshots

Rate and utilisation history, for charts. Cached ~60s.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `marketId` | path | — | Masked ID |
| `range` | query | `7d` | `1d`, `7d`, `30d` |
| `limit` | query | 200 | Capped at 2000 |

### Response — 200 OK

```json
{
  "items": [
    {
      "captured_at": "2026-08-18T03:00:00Z",
      "block": 40,
      "total_supplied": { "amount": "150000000000", "decimals": 6, "symbol": "USDC" },
      "total_borrowed": { "amount": "15000000000", "decimals": 6, "symbol": "USDC" },
      "utilization_bps": 1000,
      "supply_rate_bps": 16,
      "borrow_rate_bps": 181,
      "positions_at_risk": 0
    }
  ],
  "next_cursor": null,
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

Ordered oldest first, so it can be plotted directly.

---

## GET /accounts/{address}/summary

Everything the dashboard header needs. Cached ~5s.

| Parameter | In | Notes |
| --- | --- | --- |
| `address` | path | Wallet address, `0x…`, case-insensitive |

### Response — 200 OK

```json
{
  "address": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
  "supplied": { "amount": "50132158115", "decimals": 6, "symbol": "USDC" },
  "collateral": { "amount": "2000000000000000000", "decimals": 18, "symbol": "WETH" },
  "debt": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
  "collateral_value_usd": "580000000000",
  "debt_value_usd": "510000000000",
  "health_factor_bps": 9098,
  "health_tier": "liquidation",
  "is_liquidatable": true,
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

| Field | Notes |
| --- | --- |
| `*_value_usd` | 8 decimals — `580000000000` is `$5,800.00` |
| `health_factor_bps` | `9098` is a health factor of 0.9098. `null` when there is no debt |
| `health_tier` | `safe` (≥1.50), `caution` (1.15–1.50), `warning` (1.00–1.15), `liquidation` (<1.00) |

**An unknown address returns 200 with zeroed figures, not 404.** Every address is valid on a blockchain,
whether or not it has ever transacted — so "no position" is an empty position, not a missing resource.

Errors: `400 BAD_REQUEST` for a malformed address.

---

## GET /accounts/{address}/positions

Per-market position detail. Same fields as `summary` but one object per market, each carrying its
`market_id`, plus `max_borrowable` and `max_withdrawable_collateral`.

**Important:** these two figures are for **display only**. Before signing anything, read them from the
`PositionLens` contract directly — the backend may be a few blocks behind, and a borrow sized from a stale
figure will revert.

---

## GET /accounts/{address}/transactions

The `/history` table. Cached ~5s.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `address` | path | — | Wallet address |
| `kind` | query | all | `deposit`, `withdraw`, `borrow`, `repay`, `collateral_added`, `collateral_withdrawn`, `liquidation` |
| `from`, `to` | query | — | RFC 3339 date bounds |
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
      "block_time": "2026-08-18T03:20:11Z",
      "status": "confirmed"
    }
  ],
  "next_cursor": "b3BhcXVl",
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

Ordered newest first.

---

## GET /accounts/{address}/transactions/{transactionId}

One transaction in full, for the detail drawer. Adds the resulting health factor, gas used, and — for a
liquidation — the seized collateral and bonus. Cached ~30s.

Errors: `404 NOT_FOUND` if the transaction does not exist **or does not belong to that address**.

---

## GET /accounts/{address}/activity

The dashboard's short recent-activity list. Same item shape as `/transactions`, no cursor.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `limit` | query | 5 | Capped at 20 |

---

## GET /liquidations/eligible

Positions that can be liquidated right now. **Deliberately public and unauthenticated** — liquidators must
be able to find work before connecting a wallet. Cached ~5s, rate limited by IP.

| Parameter | In | Default | Notes |
| --- | --- | --- | --- |
| `market` | query | all | Masked market ID |
| `sort` | query | `health` | `health`, `size`, `reward` |
| `cursor`, `limit` | query | 25 | Capped at 100 |

### Response — 200 OK

```json
{
  "items": [
    {
      "borrower": "0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc",
      "market_id": "mkt_zuxjejolmpb2ynh4g32q",
      "health_factor_bps": 9098,
      "debt_to_repay": { "amount": "5100000000", "decimals": 6, "symbol": "USDC" },
      "collateral_to_seize": { "amount": "1846551724137931035", "decimals": 18, "symbol": "WETH" },
      "bonus_value_usd": "25500000000",
      "shortfall_value_usd": "0"
    }
  ],
  "next_cursor": null,
  "as_of": { "block": 45, "time": "2026-08-18T03:25:46Z" }
}
```

Two things a liquidation bot must understand:

**`shortfall_value_usd` greater than zero means the liquidation loses money.** The collateral is worth less
than the debt plus bonus, so you would pay more than you receive. Filter these out unless you have a reason
not to.

**This list can be stale.** Always call `previewLiquidation` on the contract before submitting, and expect
`PositionIsHealthy` sometimes — it means another liquidator got there first, or the price recovered.

---

## GET /liquidations/history

Completed liquidations, public. Cached ~15s. Same filters as `eligible`, plus each item carries
`liquidator`, `tx_hash`, `block_time`, and the health factor before the liquidation.

---

## GET /liquidations/{liquidationId}

One liquidation receipt: debt repaid, collateral seized, bonus, health factor before, the trigger price and
its decimals, and any shortfall. Cached ~60s.

Every field comes from the `LiquidationExecuted` event, which was designed to carry all of it — so building
a receipt never needs a follow-up call to the chain.

---

## Not in Phase 1

Deliberately absent, and why:

| Endpoint | Reason |
| --- | --- |
| `/accounts/{address}/notifications` | Needs Phase 2 auth — notification read state is off-chain |
| `/accounts/{address}/preferences` | Stores an email address. Must be guarded before it exists |
| `/accounts/{address}/email`, `/email/verify` | Same |
| `/faucet/status`, `/faucet/requests` | A rate-limited resource someone could drain on your behalf |
| `/stream` (SSE) | Live updates. Polling is sufficient for now |

---

## Trying it

```bash
curl -s http://localhost:8080/api/v1/health | python3 -m json.tool

curl -s -H "X-Request-Id: my-trace-id" -D - -o /dev/null \
  http://localhost:8080/api/v1/health

curl -s http://localhost:8080/api/v1/nope
```
