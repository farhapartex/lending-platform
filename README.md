# Lending Platform

Decentralized crypto lending and borrowing. Lend USDC to earn interest, or borrow against WETH collateral with a live safety score.

Phase 1, frontend only. Everything runs on mock data — no wallet or contract integration yet.

## Stack

Next.js 16 (App Router), React 19, TypeScript, Tailwind CSS v4.
Planned: Golang, PostgreSQL with GORM, Solidity, Docker, Kubernetes.

## Run

```bash
cd frontend
npm install
npm run dev
```

Serves on http://localhost:5173. Also available: `npm run build`, `npm run lint`, `npm run typecheck`.

