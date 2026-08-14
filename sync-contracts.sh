#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ABI_OUT="$REPO_ROOT/abi"
BACKEND_DIR="$REPO_ROOT/backend/core-service"
BINDINGS_OUT="$BACKEND_DIR/internal/chain/bindings"
FRONTEND_DIR="$REPO_ROOT/frontend"

CONTRACTS=(
  LendingPool
  CollateralVault
  LendingController
  LiquidationManager
  PositionLens
  PriceOracleAdapter
  InterestRateModel
  MockERC20
  MockAggregator
)

require() {
  command -v "$1" >/dev/null || {
    echo "missing required tool: $1" >&2
    exit 1
  }
}

require forge
require jq
require abigen

to_snake_case() {
  echo "$1" | sed 's/\([a-z0-9]\)\([A-Z]\)/\1_\2/g' | tr '[:upper:]' '[:lower:]'
}

echo "==> exporting abi from contracts"
cd "$REPO_ROOT/contracts"
forge build >/dev/null
mkdir -p "$ABI_OUT"

for contract in "${CONTRACTS[@]}"; do
  forge inspect "$contract" abi --json | jq '.' > "$ABI_OUT/$contract.json"
  printf '    %-20s %3s entries\n' "$contract" "$(jq 'length' "$ABI_OUT/$contract.json")"
done

echo "==> generating go bindings"
rm -rf "$BINDINGS_OUT"
mkdir -p "$BINDINGS_OUT"

for contract in "${CONTRACTS[@]}"; do
  abigen \
    --abi "$ABI_OUT/$contract.json" \
    --pkg bindings \
    --type "$contract" \
    --out "$BINDINGS_OUT/$(to_snake_case "$contract").go"
done

cd "$BACKEND_DIR"
gofmt -w internal/chain/bindings
go build ./...
echo "    backend build ok"

echo "==> generating frontend contract code"
cd "$FRONTEND_DIR"
npx wagmi generate >/dev/null
echo "    src/lib/contracts.generated.ts written"

echo
echo "abi       -> $ABI_OUT"
echo "backend   -> $BINDINGS_OUT"
echo "frontend  -> $FRONTEND_DIR/src/lib/contracts.generated.ts"
