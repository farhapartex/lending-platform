# Running the system

A walkthrough from an empty machine to lending, borrowing and liquidating on a local chain. Nothing here
touches real money — everything runs against a blockchain on your own laptop that you can throw away and
recreate whenever you like.

Allow about twenty minutes the first time.

[Back to the README](../README.md)

---

## What you need first

| Tool | Version used here |
| --- | --- |
| [Foundry](https://book.getfoundry.sh/getting-started/installation) | 1.7.1 |
| [Docker](https://docs.docker.com/get-docker/) with Compose | 26.x |
| [Node.js](https://nodejs.org/) | 22.x |
| [Go](https://go.dev/dl/) | 1.26 — only for running migrations |
| [MetaMask](https://metamask.io/) | any recent version |

Four terminals are handy: one for the chain, one for contracts, one for the backend, one for the web app.

---

## 1. Start the chain

```bash
anvil --host 0.0.0.0
```

Leave it running. This is a blockchain on your machine, preloaded with ten accounts holding 10,000 test ETH
each. **Everything is wiped when you stop it**, which is the point — you can break things freely.

Anvil prints ten addresses and their private keys on startup. Keep that output; you will paste from it.

> **Anvil only advances its clock when a block is mined.** On a quiet chain, time effectively stops. That
> matters later, because the price feed has a freshness window.

---

## 2. Deploy the contracts and seed test positions

In a second terminal:

```bash
cd contracts
forge build
forge script script/SeedLocal.s.sol --tc SeedLocal --rpc-url http://localhost:8545 --broadcast
```

This does the work that makes the rest of the guide possible. It deploys all seven contracts, wires them
together, mints test WETH and USDC, and then **creates ready-made positions** so the app has something to
show — including one loan that is deliberately close to the edge.

**Keep the output.** It prints every contract address, and the backend and web app both need them.

On a fresh Anvil the addresses are deterministic, so they will match these:

| Contract | Address |
| --- | --- |
| LendingPool | `0x0165878A594ca255338adfa4d48449f69242Eb8F` |
| CollateralVault | `0xa513E6E4b8f2a923D98304ec87F64353C4D5C853` |
| LendingController | `0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6` |
| LiquidationManager | `0x8A791620dd6260079BF849Dc5567aDC3F2FdC318` |
| PositionLens | `0x610178dA211FEF7D417bC0e6FeD39F05609AD788` |
| PriceOracleAdapter | `0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9` |
| InterestRateModel | `0x5FC8d32690cc91D4c39d9d3abcBD16989F875707` |
| WETH (collateral) | `0x5FbDB2315678afecb367f032d93F642f64180aa3` |
| USDC (debt) | `0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512` |
| ETH/USD feed | `0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0` |
| USDC/USD feed | `0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9` |

**If your addresses differ**, you started from a chain that already had transactions on it. Restart Anvil and
re-run the seed, or use your own addresses everywhere below.

### Who the seed script creates

The positions are set up at an ETH price of **$2,900**, and each account plays a role:

| Role | Address | Private key | Starts with |
| --- | --- | --- | --- |
| Deployer | `0xf39F…2266` | `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` | Owns the contracts |
| Lender One | `0x7099…79C8` | `0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d` | 100,000 USDC supplied |
| Lender Two | `0x3C44…93BC` | `0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a` | 50,000 USDC supplied |
| Alice — safe | `0x90F7…b906` | `0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6` | 5 WETH in, 3,000 USDC owed |
| Bob — caution | `0x15d3…6A65` | `0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a` | 3.2 WETH in, 6,900 USDC owed |
| Carol — unsafe | `0x9965…A4dc` | `0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba` | 2 WETH in, 5,100 USDC owed |
| Liquidator | `0x976E…0aa9` | `0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e` | USDC to settle loans with |

**Carol starts liquidatable** — her health factor is about 0.91. That is deliberate, so you can watch a
liquidation without having to engineer one.

> These keys are printed by every Anvil on earth. They are safe here and catastrophic anywhere else. Never
> send real funds to them.

---

## 3. Start the backend

```bash
cd backend
cp .env.example .env
docker compose up -d
```

That starts Postgres, Redis and the API on `http://localhost:8080`. Check it:

```bash
curl -s http://localhost:8080/api/v1/health
```

### Create the database tables

Migrations run from your machine, not inside the container, and the config loader reads only real
environment variables — there is no `.env` auto-loading.

```bash
cd backend/core-service
cp .env.example .env
```

Open that `.env` and paste in the contract addresses from step 2. **The migration will refuse to run while
they are blank**, because it loads the whole service config and validates the chain settings. The SQL never
uses them; it is a rough edge, not a real dependency.

```bash
set -a; . ./.env; set +a
make migrate-up
```

You should see eight migrations apply. Then restart the API so the indexer picks up the addresses:

```bash
cd backend && docker compose restart core-service
```

### Confirm the indexer is reading the chain

```bash
docker compose logs core-service | grep -i indexer
```

Look for `indexer attached to the market` and `indexed protocol events`. Within a few seconds the seeded
history should be in the database:

```bash
curl -s http://localhost:8080/api/v1/liquidations/eligible
```

That should list **Carol**. If it returns an empty list, the indexer has not caught up yet — wait a few
seconds and try again.

---

## 4. Start the web app

```bash
cd frontend
npm install
cp .env.example .env.local
```

Open `.env.local` and set:

- `NEXT_PUBLIC_CHAIN_ID=31337`
- `NEXT_PUBLIC_RPC_URL=http://localhost:8545`
- `NEXT_PUBLIC_API_BASE_URL=http://localhost:8080/api/v1`
- The contract addresses from step 2

Then:

```bash
npm run dev
```

It runs on **http://localhost:5173**, not 3000.

---

## 5. Set up MetaMask

### Add the local network

1. Open MetaMask, click the **network dropdown** at the top left
2. **Add network** → **Add a network manually**
3. Fill in:

| Field | Value |
| --- | --- |
| Network name | `Anvil Local` |
| New RPC URL | `http://localhost:8545` |
| Chain ID | `31337` |
| Currency symbol | `ETH` |
| Block explorer | leave blank |

4. Save, then switch to it

The chain ID must be exactly **31337**, matching `NEXT_PUBLIC_CHAIN_ID`. Anything else and the app will show
a wrong-network warning.

### Import a test account

1. Account menu → **Add account or hardware wallet** → **Import account**
2. Paste a private key from the table in step 2 — **Alice** is a good first choice
3. Rename it so you can tell them apart

You should see about **10,000 ETH**. The dollar value will read **$0.00** forever, because MetaMask prices
assets from its own market data and has never heard of your local chain. Watch the token amounts, not the
fiat.

### Make the tokens visible

MetaMask will not show custom tokens until you add them. On the **Tokens** tab → **Import tokens** →
**Custom token**:

| Token | Address | Decimals |
| --- | --- | --- |
| WETH | `0x5FbDB2315678afecb367f032d93F642f64180aa3` | 18 |
| USDC | `0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512` | 6 |

Symbol and decimals fill themselves in — MetaMask reads them off the contract. As Alice you should now see
**15 WETH** and **8,000 USDC**.

Note what is *not* there: the 5 WETH she already has deposited, and the 3,000 USDC she owes. Collateral sits
in the vault contract, and debt is not a token at all. Seeing that gap is the reason the dashboard exists.

> **If you restart Anvil later**, clear MetaMask's cached history or transactions will fail with
> "nonce too high": **Settings → Advanced → Clear activity tab data**, with the local network selected.

---

## 6. Connect and look around

Open `http://localhost:5173` and click through to the app.

**Connect wallet** in the top navigation → **MetaMask** → approve, choosing Alice.

You should see her address in the nav and no wrong-network banner. Now every screen describes her:

| Page | What it shows |
| --- | --- |
| **Markets** | Pool totals, both rates, utilization, the live ETH price |
| **Dashboard** | Her position, health score, and recent activity |
| **Borrow** | Her collateral and loan, with a price-drop simulator |
| **Lend** | Her supplied balance and the deposit form |
| **History** | Every transaction the indexer has recorded for her |
| **Liquidations** | Positions anyone can settle right now |

Switch MetaMask to a different imported account and watch the numbers change. Every figure is read for the
connected wallet.

---

## 7. Lend some USDC

On **Lend**, with Alice connected:

1. **Deposit** tab → type `1000`
2. **Deposit** → confirm in MetaMask
3. Watch the status: *confirm* → *submitted* → *confirmed*

Her balance goes from 0 to 1,000 USDC and the pool total rises, without a page refresh. Interest starts
accruing immediately.

Try **`0.5`** first if you want to see validation work — the minimum deposit is 1 USDC, read from the pool
contract, and it blocks before any wallet prompt.

To take it back: **Withdraw** tab → amount → **Withdraw**.

---

## 8. Borrow against collateral

On **Borrow**, still as Alice.

### Add collateral

1. **Add** tab on the left panel → type `1`
2. **Add collateral** → confirm

Her collateral rises from 5 to 6 WETH and her health score improves. If your account has no allowance yet,
you will get **two** prompts — one to approve WETH, one to deposit. Alice already has an allowance from the
seed script, so she gets one.

### Take a loan

1. **Borrow** tab on the right panel → type `1000`
2. Watch **Effect on your safety score** before you commit — it shows the health factor before and after
3. **Borrow** → confirm

The USDC lands in her wallet and her health score drops. Nothing forces you to repay on a schedule.

### See what a price fall would do

Scroll to **What if the price drops?** and drag. It re-prices her collateral and shows where liquidation
would begin. Nothing is sent anywhere — it is arithmetic on the real position.

### Repay

**Repay** tab → **All of it** → confirm. This calls the contract's own `repayAll`, which clears the interest
that accrues between reading the debt and mining the block. Paying a typed amount usually leaves dust behind.

---

## 9. Liquidate an unsafe loan

This is the part that makes the whole design work.

Carol is already liquidatable. Import the **Liquidator** account (or Lender One, who holds more USDC), switch
to it, and open **Liquidations**.

You should see Carol listed with her health factor and the bonus on offer. Click **Liquidate** →
**Repay and claim collateral** → confirm.

Within about five seconds:

- Carol disappears from the eligible list
- The **Liquidations already settled** panel below gains a row
- Your wallet is down the USDC you repaid and up the WETH you seized, plus the 5% bonus

Nothing needed reloading — the indexer saw the event and the tables refreshed.

### Making someone else liquidatable

To engineer one, drop the collateral price:

```bash
cd contracts
COLLATERAL_FEED=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0 \
DEBT_FEED=0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 \
COLLATERAL_TOKEN=0x5FbDB2315678afecb367f032d93F642f64180aa3 \
DEBT_TOKEN=0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 \
ORACLE_ADDRESS=0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9 \
COLLATERAL_PRICE=260000000000 \
DEPLOYER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
forge script script/LocalPrice.s.sol --tc LocalPrice --rpc-url http://localhost:8545 --broadcast
```

`COLLATERAL_PRICE` is in 8 decimals, so `260000000000` is $2,600. At that price Bob becomes liquidatable.

Positions are revalued every 60 seconds, so give it a minute before he appears in the list — a price move
emits no blockchain event, so nothing else would prompt a recheck.

---

## 10. Check what actually happened

The screens are one view. Here are the others.

### Through the API

```bash
curl -s http://localhost:8080/api/v1/liquidations/eligible
curl -s http://localhost:8080/api/v1/liquidations/history
curl -s "http://localhost:8080/api/v1/accounts/0x90F79bf6EB2c4f870365E785982E1f101E93b906/transactions"
```

Every response carries an `as_of` block, telling you how far the indexer has read.

### Straight from the contracts

```bash
cast call 0x610178dA211FEF7D417bC0e6FeD39F05609AD788 \
  "accountData(address)((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool))" \
  0x90F79bf6EB2c4f870365E785982E1f101E93b906 --rpc-url http://localhost:8545
```

This is the source of truth. If a screen disagrees with this, the screen is wrong.

### In the database

```bash
cd backend
docker compose exec postgres psql -U lending -d lending -c \
  "select kind, amount, block_number from user_transactions order by block_number desc limit 10;"
```

---

## When something goes wrong

**`PriceIsStale` when borrowing or liquidating.** The price feed has a one hour freshness window, and Anvil's
clock only moves when blocks are mined. Leave it idle and the price ages out. The app will tell you actions
are paused; to fix it, run the `LocalPrice` command from step 9 — without `COLLATERAL_PRICE` it restamps the
current price.

**"Nonce too high" in MetaMask.** You restarted Anvil while MetaMask still remembered the old chain.
**Settings → Advanced → Clear activity tab data**.

**Every API call returns an empty list.** The indexer is not running or has not caught up. Check
`docker compose logs core-service | grep -i indexer`, and confirm the contract addresses in
`backend/core-service/.env` match what the seed script printed.

**The wrong-network banner will not clear.** MetaMask is on a different chain ID. It must be 31337 and match
`NEXT_PUBLIC_CHAIN_ID` in `frontend/.env.local`.

**Balances look wrong after restarting Anvil.** The chain is new but MetaMask cached the old one. Clear the
activity data, and re-run the seed script.

**The browser blocks API responses.** The backend allows `http://localhost:5173` by default. Serving the app
from another port means setting `CORS_ALLOWED_ORIGINS` on the backend to match.

---

## Starting over

```bash
# stop the chain with Ctrl-C, then
cd backend && docker compose down -v   # -v also wipes the indexed data
anvil --host 0.0.0.0                   # fresh chain
```

Then repeat from step 2. Clear MetaMask's activity data as well, or it will keep trying to use the old
nonces.
