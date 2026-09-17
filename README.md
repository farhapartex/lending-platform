# FUSD

![Current system overview](Current%20system%20overview.png)

A crypto lending and borrowing app. Put money in to earn interest, or keep your ETH and borrow cash against
it.

This is a **pet project**. I am learning Solidity properly, so instead of writing a few toy contracts I am
building one full product the way a real team would — smart contracts, tests, a backend, a web app,
deployment scripts and documentation.

---

## Documentation

| Document | Who it is for |
| --- | --- |
| **[How it works](docs/how-it-works.md)** | Anyone. What the app does, what you can do with it, and the one risk borrowers must understand. No prior knowledge assumed |
| **[Running the system](docs/running-the-system.md)** | Anyone who wants to try it. From an empty machine to lending, borrowing and liquidating on a local chain, including MetaMask setup |
| **[User guide](user_guide.md)** | Anyone using the app. Every screen and every term explained in detail |
| **[Architecture](docs/architecture.md)** | Anyone reading the code. How the three pieces fit together and why |
| **[Setup](docs/setup.md)** | Anyone building from source. Tooling, tests, and the regeneration step |
| **[API reference](api_doc.md)** | Anyone building against the backend. Every endpoint, with request and response shapes |

**Start with [running the system](docs/running-the-system.md)** if you want to see it work. It is the only
document that takes you from nothing to a working setup you can click through.

---

## Where it stands

| Part | State |
| --- | --- |
| **Smart contracts** | Done. 7 contracts, 282 tests passing, including fuzz and invariant suites |
| **Deployment scripts** | Done. Deploys to a local chain with test positions ready to play with |
| **Backend** | All endpoints live and reading real data |
| **Indexer** | Reads contract events into Postgres, survives restarts, revalues positions on a timer |
| **Web app** | All 16 screens reading live data from the chain and the database |
| **Signing transactions** | Wired. Deposit, withdraw, borrow, repay, add and remove collateral, and liquidate |

Every figure on every screen is read from the chain or from indexed blockchain history. There is no
placeholder data left in the app.

### Known gaps

- **Three of the six signing flows are unverified end to end.** Collateral deposits and liquidation have been
  confirmed on-chain; borrow, repay and lend have not
- **The reorg rollback path has never run.** It is written and it deletes rows, which makes it the least
  proven code in the repository
- **No market snapshots**, so there are no historical trends to show
- **A liquidator gets no transaction of their own** — their side is recorded on the settlement, not in their
  history
- **No tests on the web app**, and the indexer has tests only for event decoding

### Deliberately not built

- One market only — lock ETH, borrow USDC. No other coins
- No insurance fund, so a very sudden crash can still leave bad debt
- No upgrades — fixing a contract means deploying a new one
- **Nobody has audited this.** It is a learning project, not somewhere to put real money

---

## Layout

```
contracts/   Solidity — the rules. The only part that can move money
backend/     Go — indexes the blockchain and serves history and totals quickly
frontend/    Next.js — the screens people use
abi/         The shared description of the contracts, generated from the Solidity
docs/        The documents linked above
```

One rule shapes all of it: **the contracts decide, everything else only displays.** The backend and the web
app can be wrong and the worst that happens is a confusing number on a screen. If a contract is wrong, money
is lost.

---

## Quick start

Four commands, assuming Foundry, Docker and Node are installed:

```bash
anvil --host 0.0.0.0                                    # terminal 1
cd contracts && forge script script/SeedLocal.s.sol \
  --tc SeedLocal --rpc-url http://localhost:8545 --broadcast   # terminal 2
cd backend && cp .env.example .env && docker compose up -d     # terminal 3
cd frontend && npm install && npm run dev                      # terminal 4
```

The database tables and MetaMask still need setting up.
**[Running the system](docs/running-the-system.md)** covers both, and explains what the seed script creates.
