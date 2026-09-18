import type { Metadata } from "next";
import {
  AssetSymbol,
  BadgeTone,
  IconName,
  SectionId,
  SectionTone,
  SurfaceElevation,
  ValueFormat,
} from "@/lib/enums";
import { formatSecondsAgo, formatValue } from "@/lib/format";
import { scaledValueToUsd, toValueScaled } from "@/lib/health";
import { formatTokenAmount } from "@/lib/token";
import { bpsToRatio } from "@/lib/units";
import {
  fusdBacking,
  fusdCollateralRatioBps,
  fusdDecimals,
  fusdOracleTimeoutSeconds,
  fusdProtocolReserve,
  fusdTotalSupply,
} from "@/content/fusdPreview";
import { backingValueScaled, collateralizationBps } from "@/lib/fusd";
import { Alert } from "@/components/ui/Alert";
import { Badge } from "@/components/ui/Badge";
import { Card } from "@/components/ui/Card";
import { PageHeader } from "@/components/ui/PageHeader";
import { Section } from "@/components/ui/Section";
import { FusdPreviewNotice } from "@/components/fusd/FusdPreviewNotice";

export const metadata: Metadata = {
  title: "FUSD backing",
  description:
    "Every FUSD in circulation, the collateral behind it, and the ratio between them. Readable without connecting a wallet.",
};

const collateralValue = backingValueScaled(fusdBacking);
const ratioBps = collateralizationBps(collateralValue, fusdTotalSupply, fusdDecimals);
const supplyValue = toValueScaled(fusdTotalSupply, fusdDecimals, 100_000_000n);
const surplus = collateralValue - supplyValue;
const isHealthy = ratioBps !== null && ratioBps >= fusdCollateralRatioBps;

export default function FusdBackingPage() {
  return (
    <>
      <PageHeader
        title="What backs FUSD"
        description="Every FUSD that exists was minted against collateral worth more than it. This page shows both sides of that claim, for anyone, without connecting a wallet."
      />

      <FusdPreviewNotice />

      <Section id={SectionId.ProtocolTotals} tone={SectionTone.Canvas}>
        <Card elevation={SurfaceElevation.Raised} className="p-6 sm:p-8">
          <dl className="grid gap-6 sm:grid-cols-4 sm:gap-8">
            <div className="flex flex-col gap-1">
              <dt className="text-sm text-ink-soft">FUSD in circulation</dt>
              <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                {formatTokenAmount(fusdTotalSupply, fusdDecimals, 0)}
              </dd>
              <p className="text-xs text-ink-faint">Claims against the protocol</p>
            </div>
            <div className="flex flex-col gap-1">
              <dt className="text-sm text-ink-soft">Collateral held</dt>
              <dd className="text-2xl font-semibold tracking-tight text-ink tabular-nums">
                {formatValue(scaledValueToUsd(collateralValue), ValueFormat.UsdCompact)}
              </dd>
              <p className="text-xs text-ink-faint">Locked in the engine</p>
            </div>
            <div className="flex flex-col gap-1">
              <dt className="text-sm text-ink-soft">Cover</dt>
              <dd className="text-2xl font-semibold tracking-tight text-brand-ink tabular-nums">
                {ratioBps === null ? "—" : formatValue(bpsToRatio(ratioBps), ValueFormat.Percent)}
              </dd>
              <p className="text-xs text-ink-faint">
                Minimum {formatValue(bpsToRatio(fusdCollateralRatioBps), ValueFormat.Percent)}
              </p>
            </div>
            <div className="flex flex-col gap-1 border-t border-line pt-4 sm:border-l sm:border-t-0 sm:pl-8 sm:pt-0">
              <dt className="text-sm text-ink-soft">Surplus</dt>
              <dd className="text-2xl font-semibold tracking-tight text-mint-ink tabular-nums">
                {formatValue(scaledValueToUsd(surplus), ValueFormat.UsdCompact)}
              </dd>
              <p className="text-xs text-ink-faint">Collateral beyond what is owed</p>
            </div>
          </dl>
        </Card>

        <Alert
          title={isHealthy ? "Every FUSD is covered" : "Cover has fallen below the minimum"}
          tone={isHealthy ? BadgeTone.Positive : BadgeTone.Critical}
          icon={isHealthy ? IconName.ShieldCheck : IconName.Warning}
          className="mt-6"
        >
          The protocol is built so the collateral it holds is always worth more than the FUSD it has issued. That is
          the first property the contract tests are written to prove, and this page is where you check it rather than
          take it on trust.
        </Alert>
      </Section>

      <Section id={SectionId.MarketSummary} tone={SectionTone.Surface} bordered>
        <div className="flex flex-col gap-1.5">
          <h2 className="text-lg font-semibold tracking-tight text-ink">Where the collateral sits</h2>
          <p className="text-sm leading-relaxed text-ink-soft">
            Each accepted asset, what the engine holds of it, and how recently its price was reported.
          </p>
        </div>

        <Card elevation={SurfaceElevation.Raised} className="mt-6 overflow-hidden p-0">
          <table className="w-full border-collapse text-left">
            <thead className="border-b border-line bg-surface-muted">
              <tr>
                <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                  Asset
                </th>
                <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                  Held
                </th>
                <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                  Value
                </th>
                <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                  Share
                </th>
                <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                  Price feed
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-line">
              {fusdBacking.map((entry) => {
                const value = toValueScaled(entry.locked, entry.decimals, entry.unitPriceScaled);
                const share = collateralValue === 0n ? 0 : Number((value * 10_000n) / collateralValue) / 10_000;
                const isFresh = entry.feedUpdatedSecondsAgo < fusdOracleTimeoutSeconds;

                return (
                  <tr key={entry.symbol}>
                    <td className="px-5 py-4">
                      <div className="flex flex-col">
                        <span className="text-sm font-medium text-ink">{entry.symbol}</span>
                        <span className="text-xs text-ink-faint">{entry.name}</span>
                      </div>
                    </td>
                    <td className="px-5 py-4 text-sm text-ink tabular-nums">
                      {formatTokenAmount(entry.locked, entry.decimals, 4)}
                    </td>
                    <td className="px-5 py-4 text-sm font-semibold text-ink tabular-nums">
                      {formatValue(scaledValueToUsd(value), ValueFormat.UsdCompact)}
                    </td>
                    <td className="px-5 py-4 text-sm text-ink-soft tabular-nums">
                      {formatValue(share, ValueFormat.Percent)}
                    </td>
                    <td className="px-5 py-4">
                      <div className="flex items-center gap-2">
                        <Badge tone={isFresh ? BadgeTone.Positive : BadgeTone.Caution}>
                          {isFresh ? "Fresh" : "Stale"}
                        </Badge>
                        <span className="text-xs text-ink-faint">
                          {formatSecondsAgo(entry.feedUpdatedSecondsAgo)}
                        </span>
                      </div>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </Card>

        <div className="mt-6 grid gap-6 sm:grid-cols-2">
          <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-2 p-6">
            <h3 className="text-sm font-medium text-ink">Protocol reserve</h3>
            <p className="text-xl font-semibold tracking-tight text-ink tabular-nums">
              {formatTokenAmount(fusdProtocolReserve, fusdDecimals, 2)} {AssetSymbol.Fusd}
            </p>
            <p className="text-sm leading-relaxed text-ink-soft">
              Built from the two percent taken on each liquidation. It sits outside the backing above and is not
              counted toward cover.
            </p>
          </Card>

          <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-2 p-6">
            <h3 className="text-sm font-medium text-ink">If a feed goes quiet</h3>
            <p className="text-sm leading-relaxed text-ink-soft">
              A price older than three hours is treated as unusable. The protocol stops minting and stops releasing
              collateral rather than pricing either off a number nobody can verify. Repaying always keeps working.
            </p>
          </Card>
        </div>
      </Section>
    </>
  );
}
