import type { Metadata } from "next";
import {
  AssetSymbol,
  BadgeTone,
  ButtonSize,
  ButtonVariant,
  IconName,
  SectionId,
  SectionTone,
  SurfaceElevation,
  ValueFormat,
  WalletGatePurpose,
} from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { formatHealthFactor, scaledValueToUsd, toValueScaled } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { bpsToRatio } from "@/lib/units";
import {
  fusdDecimals,
  fusdLiquidationThresholdBps,
  fusdLiquidatorBonusBps,
  fusdProtocolFeeBps,
  fusdUnsafePositions,
} from "@/content/fusdPreview";
import { fusdHealthFactorBps, fusdHealthTier, liquidationRewardScaled, positionCollateralValueScaled } from "@/lib/fusd";
import { Alert } from "@/components/ui/Alert";
import { AddressDisplay } from "@/components/ui/AddressDisplay";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { WalletGate } from "@/components/app/WalletGate";
import { HealthBadge } from "@/components/borrow/HealthBadge";
import { FusdPreviewNotice } from "@/components/fusd/FusdPreviewNotice";

export const metadata: Metadata = {
  title: "FUSD liquidations",
  description:
    "Positions whose health factor has fallen below 1. Repay part or all of the debt and take the collateral, plus the published bonus.",
};

const rows = fusdUnsafePositions.map((position) => {
  const collateralValue = positionCollateralValueScaled(position.collateral);
  const debtValue = toValueScaled(position.minted, fusdDecimals, 100_000_000n);
  const factorBps = fusdHealthFactorBps(collateralValue, position.minted, fusdDecimals, fusdLiquidationThresholdBps);

  return {
    ...position,
    collateralValue,
    debtValue,
    factorBps,
    tier: fusdHealthTier(factorBps),
    bonusValue: liquidationRewardScaled(debtValue, fusdLiquidatorBonusBps),
  };
});

export default function FusdLiquidationsPage() {
  return (
    <>
      <PageHeader
        title="FUSD liquidations"
        description="When a position falls below a health factor of 1, anyone can repay its debt and take the collateral behind it. You keep a published bonus for doing it, and the protocol keeps a smaller share."
      />

      <FusdPreviewNotice />

      <Section id={SectionId.LiquidationsList} tone={SectionTone.Canvas}>
        <div className="flex flex-col gap-6">
          <div className="grid gap-6 sm:grid-cols-3">
            <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-1 p-6">
              <span className="text-sm text-ink-soft">You keep</span>
              <span className="text-2xl font-semibold tracking-tight text-mint-ink tabular-nums">
                {formatValue(bpsToRatio(fusdLiquidatorBonusBps), ValueFormat.Percent)}
              </span>
              <span className="text-xs text-ink-faint">On top of the debt you repay</span>
            </Card>
            <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-1 p-6">
              <span className="text-sm text-ink-soft">Protocol keeps</span>
              <span className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                {formatValue(bpsToRatio(fusdProtocolFeeBps), ValueFormat.Percent)}
              </span>
              <span className="text-xs text-ink-faint">Into the reserve</span>
            </Card>
            <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-1 p-6">
              <span className="text-sm text-ink-soft">You may repay</span>
              <span className="text-2xl font-semibold tracking-tight text-ink">Part or all</span>
              <span className="text-xs text-ink-faint">Unlike the lending market</span>
            </Card>
          </div>

          <Alert title="A partial close has to leave them healthier" tone={BadgeTone.Neutral} icon={IconName.Info}>
            You do not have to clear the whole debt. The engine checks that the position ends with a higher health
            factor than it started, and reverts if it does not, so a badly sized partial cannot make things worse.
          </Alert>

          {rows.length === 0 ? (
            <EmptyState
              title="No positions are eligible right now"
              description="Every FUSD position is above a health factor of 1. This list fills up when collateral prices fall."
              icon={IconName.ShieldCheck}
            />
          ) : (
            <WalletGate purpose={WalletGatePurpose.Liquidate}>
              <Card elevation={SurfaceElevation.Raised} className="overflow-hidden p-0">
                <table className="w-full border-collapse text-left">
                  <thead className="border-b border-line bg-surface-muted">
                    <tr>
                      <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                        Borrower
                      </th>
                      <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                        Collateral
                      </th>
                      <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                        FUSD owed
                      </th>
                      <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                        Health
                      </th>
                      <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                        Your bonus
                      </th>
                      <th scope="col" className="px-5 py-3" />
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-line">
                    {rows.map((row) => (
                      <tr key={row.id}>
                        <td className="px-5 py-4">
                          <AddressDisplay address={row.borrower} />
                        </td>
                        <td className="px-5 py-4">
                          <div className="flex flex-col gap-0.5">
                            {row.collateral.map((entry) => (
                              <span key={entry.symbol} className="text-sm text-ink tabular-nums">
                                {formatTokenAmount(entry.amount, entry.decimals, 4)} {entry.symbol}
                              </span>
                            ))}
                            <span className="text-xs text-ink-faint tabular-nums">
                              {formatValue(scaledValueToUsd(row.collateralValue), ValueFormat.UsdPrice)}
                            </span>
                          </div>
                        </td>
                        <td className="px-5 py-4 text-sm font-semibold text-ink tabular-nums">
                          {formatTokenAmount(row.minted, fusdDecimals, 2)} {AssetSymbol.Fusd}
                        </td>
                        <td className="px-5 py-4">
                          <div className="flex items-center gap-2">
                            <HealthBadge tier={row.tier} />
                            <span className="text-sm text-ink-soft tabular-nums">{formatHealthFactor(row.factorBps, 4)}</span>
                          </div>
                        </td>
                        <td className="px-5 py-4 text-sm font-semibold text-mint-ink tabular-nums">
                          {formatValue(scaledValueToUsd(row.bonusValue), ValueFormat.UsdPrice)}
                        </td>
                        <td className="px-5 py-4 text-right">
                          <Button size={ButtonSize.Sm} variant={ButtonVariant.Secondary} disabled>
                            Settle
                          </Button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </Card>
            </WalletGate>
          )}

          <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-2 p-6 sm:p-7">
            <h2 className="text-base font-semibold tracking-tight text-ink">Which collateral you receive</h2>
            <p className="text-sm leading-relaxed text-ink-soft">
              A position can be backed by more than one asset. When you settle part of a debt you take a share of what
              is there, so a position holding both WETH and WBTC pays you in both rather than letting you pick the one
              you prefer.
            </p>
          </Card>
        </div>
      </Section>
    </>
  );
}
