# Architecture

How the three pieces fit together and why they are split the way they are.

[Back to the README](../README.md)

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
