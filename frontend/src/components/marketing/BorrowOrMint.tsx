"use client";

import { AppRoute, BadgeTone, ButtonSize, ButtonVariant, IconName, SectionId, SectionTone, SurfaceElevation, ValueFormat } from "@/lib/enums";
import { formatValue } from "@/lib/format";
import { bpsToRatio } from "@/lib/units";
import { borrowOrMintContent } from "@/content/landing";
import { useMarketData } from "@/hooks/useMarketData";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { Skeleton } from "@/components/ui/Skeleton";

type Term = {
  label: string;
  value: string;
};

function TermList({ terms }: { terms: Term[] }) {
  return (
    <dl className="mt-6 divide-y divide-line border-t border-line">
      {terms.map((term) => (
        <div key={term.label} className="flex items-baseline justify-between gap-4 py-3">
          <dt className="text-sm text-ink-soft">{term.label}</dt>
          <dd className="text-sm font-semibold text-ink tabular-nums">{term.value}</dd>
        </div>
      ))}
    </dl>
  );
}

export function BorrowOrMint() {
  const { data } = useMarketData();

  const marketTerms: Term[] =
    data === undefined
      ? []
      : [
          { label: "What it costs", value: `${formatValue(bpsToRatio(data.borrowAprBps), ValueFormat.Percent)} a year` },
          { label: "Borrow up to", value: formatValue(bpsToRatio(data.maxLtvBps), ValueFormat.Percent) },
          { label: "Liquidated at", value: formatValue(bpsToRatio(data.liquidationThresholdBps), ValueFormat.Percent) },
          { label: "Liquidation penalty", value: formatValue(bpsToRatio(data.liquidationBonusBps), ValueFormat.Percent) },
          { label: "Limited by", value: "Pool liquidity" },
        ];

  const fusdTerms: Term[] = [
    { label: "What it costs", value: "Nothing" },
    { label: "Borrow up to", value: "66.67%" },
    { label: "Liquidated at", value: "Health factor 1.00" },
    { label: "Liquidation penalty", value: "8% + 2%" },
    { label: "Limited by", value: "Your collateral only" },
  ];

  return (
    <Section id={SectionId.BorrowOrMint} tone={SectionTone.Canvas}>
      <SectionHeading
        sectionId={SectionId.BorrowOrMint}
        eyebrow={borrowOrMintContent.eyebrow}
        title={borrowOrMintContent.title}
        description={borrowOrMintContent.description}
      />

      <div className="mt-10 grid gap-6 lg:grid-cols-2 lg:gap-8">
        <Card elevation={SurfaceElevation.Raised} className="flex flex-col p-6 sm:p-7">
          <div className="flex flex-wrap items-center justify-between gap-3">
            <h3 className="text-lg font-semibold tracking-tight text-ink">Borrow USDC</h3>
            <Badge tone={BadgeTone.Positive}>Live now</Badge>
          </div>

          <p className="mt-3 text-sm leading-relaxed text-ink-soft">
            Take USDC that other people have deposited into the shared pool. You pay them interest for it, and how
            much you can take depends on how much is sitting there unlent.
          </p>

          {data === undefined ? (
            <div className="mt-6 flex flex-col gap-3">
              <Skeleton className="h-10 w-full" />
              <Skeleton className="h-10 w-full" />
              <Skeleton className="h-10 w-full" />
            </div>
          ) : (
            <TermList terms={marketTerms} />
          )}

          <p className="mt-6 text-sm leading-relaxed text-ink-soft">
            Suits you if you want the most borrowing power against your collateral and do not mind paying for it.
          </p>

          <div className="mt-6">
            <Button href={AppRoute.Borrow} size={ButtonSize.Md} trailingIcon={IconName.ArrowRight}>
              Borrow against WETH
            </Button>
          </div>
        </Card>

        <Card elevation={SurfaceElevation.Flat} className="flex flex-col border-dashed p-6 sm:p-7">
          <div className="flex flex-wrap items-center justify-between gap-3">
            <h3 className="text-lg font-semibold tracking-tight text-ink">Mint FUSD</h3>
            <Badge tone={BadgeTone.Caution}>In development</Badge>
          </div>

          <p className="mt-3 text-sm leading-relaxed text-ink-soft">
            FUSD is a dollar the protocol creates the moment you ask for it, backed by the collateral you lock. There
            is no pool to run dry and nobody to pay interest to, because nobody lent you anything.
          </p>

          <TermList terms={fusdTerms} />

          <p className="mt-6 text-sm leading-relaxed text-ink-soft">
            Suits you if you want to hold the dollars for a long time and care more about what it costs than about
            squeezing out the largest possible loan.
          </p>

          <div className="mt-6">
            <Button href={AppRoute.Fusd} size={ButtonSize.Md} variant={ButtonVariant.Secondary} trailingIcon={IconName.ArrowRight}>
              Read how it will work
            </Button>
          </div>
        </Card>
      </div>

      <p className="mt-6 flex items-start gap-2 text-sm text-ink-faint">
        <Icon name={IconName.Info} className="mt-0.5 size-4" />
        {borrowOrMintContent.footnote}
      </p>
    </Section>
  );
}
