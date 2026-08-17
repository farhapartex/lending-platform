# Lending Platform

A crypto lending and borrowing app. Put money in to earn interest, or keep your ETH and borrow cash against it.

This is a **pet project**. I am learning Solidity properly, so instead of writing a few toy contracts I am
building one full product the way a real team would — smart contracts, tests, a backend, a web app,
deployment scripts and documentation.

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
| **Backend** | Just started — config and blockchain connection are in |
| **Web app** | All 16 screens built, still showing example numbers rather than live ones |

The next job is connecting the three pieces together so real money moves through real screens.

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

## Running it

### Contracts

```bash
cd contracts
forge test
```

To get a local chain with money and a few ready-made loans, including one that is unsafe on purpose:

```bash
anvil
forge script script/SeedLocal.s.sol --tc SeedLocal --rpc-url http://localhost:8545 --broadcast
```

It prints every address you need. `script/LocalPrice.s.sol` moves the ETH price so you can watch a loan go
from safe to unsafe.

### Web app

```bash
cd frontend
npm install
npm run dev
```

Runs on http://localhost:5173. Also `npm run build`, `npm run lint`, `npm run typecheck`.

### Backend

```bash
cd backend
docker compose up -d

cd core-service
go test ./...
go run ./cmd/api
```

### Regenerating contract interfaces

After changing any contract:

```bash
./sync-contracts.sh
```

That rebuilds the contracts and regenerates the Go and TypeScript code, so the backend and web app can
never drift out of step with the Solidity.
