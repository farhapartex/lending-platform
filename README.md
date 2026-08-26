# Lending Platform

A crypto lending and borrowing app. Put money in to earn interest, or keep your ETH and borrow cash against it.

This is a **pet project**. I am learning Solidity properly, so instead of writing a few toy contracts I am
building one full product the way a real team would — smart contracts, tests, a backend, a web app,
deployment scripts and documentation.

📖 **New here and just want to use it?** Read the **[user guide](user_guide.md)** instead. It explains every
screen and every term from scratch. The rest of this file is about building and running the project.

---

## What problem is this solving?

Say you own some ETH. You believe it will be worth more later, so you do not want to sell it. But you need
cash today.

A bank could lend you money against it, but the bank wants to know who you are, check your credit, and
approve you first. That takes days, and plenty of people get refused.

This app does the same job with no bank in the middle. You lock up your ETH, and you can immediately borrow
dollars against it. Nobody checks your name or your credit score. There is no application form and no
waiting.

The trick is simple: **you always lock up more than you borrow.** That is what replaces the credit check.

---

## What you can actually do

There are two kinds of people using it.

### If you want to earn

You have some USDC (a digital dollar) sitting idle.

1. Connect your wallet
2. Deposit your USDC into the shared pool
3. You start earning interest straight away, paid by the people borrowing it
4. Take your money out whenever you like, plus whatever it earned

No lock-in period, no minimum term.

### If you want to borrow

You own some ETH and want cash without selling it.

1. Connect your wallet
2. Lock up your ETH as **collateral** — it stays yours, it is just held while you owe money
3. Borrow up to **75%** of what your ETH is worth. Lock $10,000 of ETH, borrow up to $7,500

> Strictly, the collateral is **WETH** — the token version of ETH, worth exactly the same. Plain ETH cannot
> be deposited because the contracts work with tokens. The [user guide](user_guide.md) explains why, and what
> to do if all you hold is ETH.
4. Pay it back whenever you want. Interest builds up every second, so paying back sooner costs less
5. Once you have repaid, your ETH is released

Nobody sets a deadline. You can keep the loan open for years if you keep it safe.

---

## The one thing borrowers must understand

Your ETH can drop in value. If it drops far enough, your loan becomes too big compared to what is backing
it — and then anyone in the world is allowed to pay off your loan and take your collateral instead.

That is called **liquidation**, and it is the part of these apps that catches people out. So the app shows
you one number, all the time:

| Health score | What it means |
| --- | --- |
| Above 1.50 | Comfortable |
| 1.15 to 1.50 | Getting tight, keep an eye on it |
| Just above 1.00 | Danger |
| Below 1.00 | Anyone can close your loan |

Two ways to push that number back up: add more collateral, or pay back some of what you owe.

You are allowed to borrow up to 75% of your collateral, but you are only liquidated at 80%. That gap is
deliberate breathing room — without it, anyone who borrowed the maximum would be one small price wobble away
from losing their collateral.

One more thing worth knowing: **the health score falls slowly on its own**, even if the ETH price never
moves, because interest keeps adding to what you owe. A loan you opened months ago and forgot about is
riskier than you left it.

---

## Why liquidation is a thing at all

It sounds harsh, so it is worth explaining.

Everyone lending money into the pool needs to get paid back. If a loan is worth more than the collateral
behind it, that money is simply gone, and the lenders lose out.

So the app pays strangers to prevent that. Anyone who closes an unsafe loan gets the collateral at a **5%
discount** as their reward. It is not charity — it is a job, and it protects the people who deposited.

It also means nobody has to be watching. There is no company monitoring loans. The reward is enough to make
people watch on their own.

---

## Where it stands right now

| Part | State |
| --- | --- |
| **Smart contracts** | Done. 7 contracts, 282 tests passing |
| **Deployment scripts** | Done. Runs on a local chain with test positions ready to play with |
| **Contract interfaces** for the backend and web app | Generated |
| **Backend** | 6 of 7 read endpoints live, all fully tested |
| **Indexer** | Not started. This is the next job |
| **Web app** | All 16 screens built. Three panels read live data; the rest still show example numbers |
| **Signing transactions** | Not wired yet. Action screens open and validate, but cannot submit |

**The indexer is the next piece of work.** Nothing currently writes blockchain history into the database, so
every endpoint answers correctly but with no rows. The screens that read it say so honestly rather than
pretending the history is empty.

After that: connecting the action buttons to the contracts, so money actually moves through real screens.

### What is deliberately not built yet

- One market only — lock ETH, borrow USDC. No other coins
- No insurance fund, so a very sudden crash can still leave bad debt
- No upgrades — fixing a contract means deploying a new one
- Nobody has audited this. It is a learning project, not somewhere to put real money

---

## How it is put together

```
contracts/   Solidity — the rules. The only part that can move money
backend/     Go — reads the blockchain and serves history and totals quickly
frontend/    Next.js — the screens people use
abi/         The shared description of the contracts, generated from the Solidity
```

One rule shapes all of it: **the contracts decide, everything else only displays.** The backend and the web
app can be wrong and the worst that happens is a confusing number on a screen. If a contract is wrong,
money is lost. So the contracts get the tests, and every figure shown before you sign anything is read
straight from them.

### Stack

- **Contracts** — Solidity, Foundry, OpenZeppelin
- **Backend** — Go, PostgreSQL, GORM, Redis, go-ethereum
- **Web app** — Next.js 16, React 19, TypeScript, Tailwind CSS v4, wagmi, viem
- **Infra** — Docker, docker-compose

---

## Setting it up

Three pieces, and they stack in this order: **contracts → backend → web app**. The contracts publish the
addresses that the other two need, so start there.

### What you need installed

| Tool | Version used here | For |
| --- | --- | --- |
| [Foundry](https://book.getfoundry.sh/getting-started/installation) | 1.7.1 | Contracts, tests, local chain |
| [Go](https://go.dev/dl/) | 1.26 | Backend |
| [Docker](https://docs.docker.com/get-docker/) with Compose | 26.x | Postgres and Redis |
| [Node.js](https://nodejs.org/) | 22.x | Web app |

For the abi regeneration step you also need `abigen` (from go-ethereum) and `jq`. Everything else runs
without them.

---

### 1. Contracts

```bash
cd contracts
forge build
forge test
```

You should see **282 tests passing**. That includes fuzz tests and stateful invariant tests, so the first
run takes about 20 seconds.

#### A local chain to play with

In one terminal, start a local blockchain and leave it running:

```bash
anvil
```

In another, deploy the contracts and create some ready-made positions — including one that is deliberately
unsafe:

```bash
cd contracts
forge script script/SeedLocal.s.sol --tc SeedLocal --rpc-url http://localhost:8545 --broadcast
```

**Keep the output.** It prints every contract address, and the backend and web app both need them.

To deploy without the test positions, use `Deploy.s.sol --tc Deploy` and fill in `contracts/.env` first
(copy `contracts/.env.example`).

#### Moving the price

`script/LocalPrice.s.sol` changes the collateral price on the local chain, which is how you push a healthy
loan into liquidation territory and watch what happens:

```bash
forge script script/LocalPrice.s.sol --tc LocalPrice --rpc-url http://localhost:8545 --broadcast
```

> **If you leave anvil running for hours and then get a `PriceIsStale` error**, the machine clock has drifted
> past the price feed's freshness window. Re-run `LocalPrice.s.sol` to stamp a fresh price.

---

### 2. Backend

```bash
cd backend
cp .env.example .env
docker compose up -d
```

That starts **Postgres**, **Redis** and the **API** on `http://localhost:8080`. Check it:

```bash
curl -s http://localhost:8080/api/v1/health
```

> **If Redis fails to start with "port is already allocated"**, something else on your machine is using
> 6379. The port is configurable, so no file needs editing:
>
> ```bash
> REDIS_PORT=6380 docker compose up -d
> ```
>
> Only the host port changes. The API still reaches Redis internally on 6379.

#### Creating the database tables

Migrations run from your machine, not from inside the container, and **the config loader reads only real
environment variables — there is no `.env` auto-loading.** So set up an env file and source it:

```bash
cd backend/core-service
cp .env.example .env
```

Open `.env` and paste in the contract addresses that `SeedLocal` printed. Then:

```bash
set -a; . ./.env; set +a
make migrate-up
```

Other migration commands: `make migrate-status`, `make migrate-down`, `make migrate-reset`.

> **Why does a database migration need contract addresses?** Because the migrate command loads the whole
> service config, which validates the chain settings. It is a rough edge, not a real dependency — the
> addresses are never used by the SQL.

#### Running the API outside Docker

Useful when you want fast rebuilds:

```bash
cd backend/core-service
set -a; . ./.env; set +a
make run
```

#### Tests

The repository tests run against **real Postgres**, and the cache tests against **real Redis**, so both must
be up:

```bash
cd backend
docker compose up -d postgres redis

cd core-service
export DATABASE_URL="postgres://lending:lending@localhost:5432/lending?sslmode=disable"
make test
```

Each test runs in a transaction that is rolled back, so nothing is left behind.

---

### 3. Web app

```bash
cd frontend
npm install
cp .env.example .env.local
```

Open `.env.local` and fill in:

- `NEXT_PUBLIC_CHAIN_ID` — `31337` for local anvil, `11155111` for Sepolia
- `NEXT_PUBLIC_RPC_URL` — `http://localhost:8545` for local
- The contract addresses from `SeedLocal`
- `NEXT_PUBLIC_API_BASE_URL` — `http://localhost:8080/api/v1`

Then:

```bash
npm run dev
```

Runs on **http://localhost:5173** — not 3000.

Other commands: `npm run build`, `npm run lint`, `npm run typecheck`.

> The backend allows browser requests from `http://localhost:5173` by default. If you serve the app from a
> different port, set `CORS_ALLOWED_ORIGINS` on the backend to match, or the browser will block every
> response.

---

### Regenerating contract interfaces

After changing any contract:

```bash
./sync-contracts.sh
```

That rebuilds the contracts and regenerates the Go bindings and the TypeScript hooks from one shared set of
abi files, so the backend and web app can never drift out of step with the Solidity.

---

## Documentation

| Document | Who it is for |
| --- | --- |
| **[User guide](user_guide.md)** | Anyone using the app. Explains every page, every term, and what you can and cannot do — assuming no prior knowledge beyond having used MetaMask |
| **[API reference](api_doc.md)** | Anyone building against the backend. Every endpoint, with request and response shapes |

The user guide is the place to start if you want to understand *what this thing does* rather than how it is
built. It covers what an ERC-20 token is, why your ETH cannot be deposited directly, how the health factor
works, and what happens during a liquidation.
