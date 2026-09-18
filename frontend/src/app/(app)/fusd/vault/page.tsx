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
import { formatHealthFactor, scaledValueToUsd } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { bpsToRatio } from "@/lib/units";
import {
  fusdCollaterals,
  fusdDecimals,
  fusdLiquidationThresholdBps,
  fusdLiquidatorBonusBps,
  fusdMinted,
  fusdProtocolFeeBps,
  fusdWalletBalance,
} from "@/content/fusdPreview";
import { fusdHealthFactorBps, fusdHealthTier, mintCapacity, totalCollateralValueScaled } from "@/lib/fusd";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { MetricRow } from "@/components/ui/MetricRow";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { WalletGate } from "@/components/app/WalletGate";
import { HealthBadge } from "@/components/borrow/HealthBadge";
import { HealthScoreGauge } from "@/components/borrow/HealthScoreGauge";
import { CollateralTable } from "@/components/fusd/CollateralTable";
import { FusdPreviewNotice } from "@/components/fusd/FusdPreviewNotice";

export const metadata: Metadata = {
  title: "FUSD vault",
  description:
    "Lock WETH or WBTC, mint FUSD against it, and watch one health factor covering every asset you have locked.",
};

const collateralValue = totalCollateralValueScaled(fusdCollaterals);
const factorBps = fusdHealthFactorBps(collateralValue, fusdMinted, fusdDecimals, fusdLiquidationThresholdBps);
const tier = fusdHealthTier(factorBps);
const capacity = mintCapacity(collateralValue, fusdMinted, fusdDecimals, fusdLiquidationThresholdBps);

export default function FusdVaultPage() {
  return (
    <>
      <PageHeader
        title="Your FUSD vault"
        description="Lock collateral, mint dollars against it, and repay whenever you choose. One health factor covers everything you have locked, whichever assets they are."
      />

      <FusdPreviewNotice />

      <Section id={SectionId.DashboardOverview} tone={SectionTone.Canvas}>
        <WalletGate purpose={WalletGatePurpose.PersonalData}>
          <div className="flex flex-col gap-6">
            <Card elevation={SurfaceElevation.Raised} className="p-6 sm:p-7">
              <dl className="grid gap-6 sm:grid-cols-4">
                <div className="flex flex-col gap-1">
                  <dt className="text-sm text-ink-soft">Collateral locked</dt>
                  <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                    {formatValue(scaledValueToUsd(collateralValue), ValueFormat.UsdPrice)}
                  </dd>
                  <p className="text-xs text-ink-faint">Across every asset</p>
                </div>
                <div className="flex flex-col gap-1">
                  <dt className="text-sm text-ink-soft">FUSD minted</dt>
                  <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                    {formatTokenAmount(fusdMinted, fusdDecimals, 2)}
                  </dd>
                  <p className="text-xs text-ink-faint">What you owe back</p>
                </div>
                <div className="flex flex-col gap-1">
                  <dt className="text-sm text-ink-soft">Still mintable</dt>
                  <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                    {formatTokenAmount(capacity, fusdDecimals, 2)}
                  </dd>
                  <p className="text-xs text-ink-faint">Before you reach the limit</p>
                </div>
                <div className="flex flex-col gap-1 border-t border-line pt-4 sm:border-l sm:border-t-0 sm:pl-6 sm:pt-0">
                  <dt className="text-sm text-ink-soft">Health factor</dt>
                  <dd className="flex items-center gap-2">
                    <span className="text-2xl font-semibold tracking-tight text-brand-ink tabular-nums">
                      {formatHealthFactor(factorBps)}
                    </span>
                    <HealthBadge tier={tier} />
                  </dd>
                  <p className="text-xs text-ink-faint">Liquidation below 1.00</p>
                </div>
              </dl>
            </Card>

            <Card className="grid gap-8 p-6 sm:p-7 lg:grid-cols-[minmax(0,1fr)_minmax(0,20rem)]">
              <HealthScoreGauge factorBps={factorBps} tier={tier} />
              <div className="border-t border-line pt-6 lg:border-l lg:border-t-0 lg:pl-8 lg:pt-0">
                <h3 className="text-sm font-medium text-ink">The rules</h3>
                <dl className="mt-3 divide-y divide-line">
                  <MetricRow
                    label="Mint up to"
                    value={formatValue(bpsToRatio(fusdLiquidationThresholdBps), ValueFormat.Percent)}
                    hint="Of your collateral value"
                  />
                  <MetricRow label="Liquidated below" value="1.00" hint="Health factor, not a percentage" />
                  <MetricRow
                    label="Liquidator takes"
                    value={formatValue(bpsToRatio(fusdLiquidatorBonusBps), ValueFormat.Percent)}
                    hint="On top of the debt they repay"
                  />
                  <MetricRow
                    label="Protocol keeps"
                    value={formatValue(bpsToRatio(fusdProtocolFeeBps), ValueFormat.Percent)}
                    hint="Into the reserve"
                  />
                  <MetricRow label="Cost to borrow" value="None" hint="No interest in version one" />
                </dl>
              </div>
            </Card>

            <div className="flex flex-col gap-4">
              <div className="flex flex-col gap-1.5">
                <h2 className="text-lg font-semibold tracking-tight text-ink">Your collateral</h2>
                <p className="text-sm leading-relaxed text-ink-soft">
                  Every asset you have locked counts toward the same health factor. Unlocking one affects all of it.
                </p>
              </div>
              <CollateralTable collaterals={fusdCollaterals} />
            </div>

            <div className="grid gap-6 lg:grid-cols-2">
              <Card elevation={SurfaceElevation.Raised} className="flex flex-col gap-5 p-6 sm:p-7">
                <div className="flex flex-col gap-1.5">
                  <h2 className="text-lg font-semibold tracking-tight text-ink">Mint FUSD</h2>
                  <p className="text-sm leading-relaxed text-ink-soft">
                    Create dollars against what you have locked. You can lock collateral and mint in a single
                    transaction.
                  </p>
                </div>

                <dl className="divide-y divide-line border-y border-line">
                  <MetricRow
                    label="Available to mint"
                    value={`${formatTokenAmount(capacity, fusdDecimals, 2)} ${AssetSymbol.Fusd}`}
                    hint="Based on the collateral already locked"
                  />
                </dl>

                <Button size={ButtonSize.Lg} fullWidth disabled>
                  Mint FUSD
                </Button>
              </Card>

              <Card elevation={SurfaceElevation.Raised} className="flex flex-col gap-5 p-6 sm:p-7">
                <div className="flex flex-col gap-1.5">
                  <h2 className="text-lg font-semibold tracking-tight text-ink">Repay and unlock</h2>
                  <p className="text-sm leading-relaxed text-ink-soft">
                    Send FUSD back to reduce what you owe. Clear it entirely and every asset is free to withdraw.
                  </p>
                </div>

                <dl className="divide-y divide-line border-y border-line">
                  <MetricRow
                    label="FUSD in your wallet"
                    value={`${formatTokenAmount(fusdWalletBalance, fusdDecimals, 2)} ${AssetSymbol.Fusd}`}
                    hint="What you can repay with right now"
                  />
                  <MetricRow
                    label="Outstanding"
                    value={`${formatTokenAmount(fusdMinted, fusdDecimals, 2)} ${AssetSymbol.Fusd}`}
                    hint="Repay all of it to unlock everything"
                  />
                </dl>

                <Button size={ButtonSize.Lg} fullWidth variant={ButtonVariant.Secondary} disabled>
                  Repay FUSD
                </Button>
              </Card>
            </div>

            <Alert title="Two health factors, one wallet" tone={BadgeTone.Caution} icon={IconName.Info}>
              Collateral locked here is separate from collateral in the lending market, and the two are judged by
              different rules. A safe position on one side says nothing about the other.
            </Alert>
          </div>
        </WalletGate>
      </Section>
    </>
  );
}
