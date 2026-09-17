# Setting up a development environment

Building and running the project from source. If you only want to use the app, read
[running the system](running-the-system.md) instead, which walks through the whole thing end to end.

[Back to the README](../README.md)

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

It reads the feed and token addresses from the environment, so they have to be supplied:

```bash
COLLATERAL_FEED=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0 \
DEBT_FEED=0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 \
COLLATERAL_TOKEN=0x5FbDB2315678afecb367f032d93F642f64180aa3 \
DEBT_TOKEN=0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 \
ORACLE_ADDRESS=0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9 \
DEPLOYER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
forge script script/LocalPrice.s.sol --tc LocalPrice --rpc-url http://localhost:8545 --broadcast
```

Add `COLLATERAL_PRICE=260000000000` to set a specific price — the value is in 8 decimals, so that is $2,600.
Without it the current price is restamped with a fresh timestamp.

> **`PriceIsStale` errors** mean the feed has aged past its one hour window. Anvil only advances its clock
> when a block is mined, so an idle chain will drift into this. The command above clears it.

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
