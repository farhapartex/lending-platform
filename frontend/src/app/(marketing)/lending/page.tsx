import type { Metadata } from "next";
import {
  AppRoute,
  BadgeTone,
  ButtonSize,
  ButtonVariant,
  IconName,
  SectionId,
  SectionTone,
  SurfaceElevation,
} from "@/lib/enums";
import { lendingLandingContent, lendingPaths, lendingRiskContent } from "@/content/lendingLanding";
import { feeItems } from "@/content/protocol";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { MetricRow } from "@/components/ui/MetricRow";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";
import { LiveMarketStrip } from "@/components/marketing/LiveMarketStrip";
import { ClosingCta } from "@/components/marketing/ClosingCta";

export const metadata: Metadata = {
  title: "Lending",
  description:
    "Deposit USDC to earn interest paid by borrowers, or lock WETH and borrow against it without selling. Rates set by pool utilization.",
};

export default function LendingPage() {
  return (
    <>
      <section
        id={SectionId.Hero}
        aria-labelledby={`${SectionId.Hero}-heading`}
        className="relative overflow-hidden border-b border-line bg-gradient-to-b from-accent-soft/60 via-canvas to-canvas"
      >
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -top-40 left-1/3 size-[38rem] rounded-full bg-accent-muted/35 blur-3xl"
        />

        <Container className="relative flex flex-col items-start gap-6 py-16 sm:py-20 lg:py-24">
          <Badge tone={BadgeTone.Positive}>{lendingLandingContent.eyebrow}</Badge>

          <h1
            id={`${SectionId.Hero}-heading`}
            className="max-w-3xl text-balance text-4xl font-semibold leading-[1.1] tracking-tight text-ink sm:text-5xl"
          >
            {lendingLandingContent.title}
          </h1>

          <p className="max-w-2xl text-pretty text-lg leading-relaxed text-ink-soft">
            {lendingLandingContent.description}
          </p>

          <Button href={AppRoute.Login} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight}>
            Connect a wallet
          </Button>
        </Container>
      </section>

      <Section id={SectionId.ProtocolStats} tone={SectionTone.Canvas}>
        <LiveMarketStrip />
        <p className="mt-4 flex items-start gap-2 text-sm text-ink-faint">
          <Icon name={IconName.Info} className="mt-0.5 size-4" />
          {lendingLandingContent.note}
        </p>
      </Section>

      <Section id={SectionId.ValueProps} tone={SectionTone.Surface} bordered>
        <SectionHeading
          sectionId={SectionId.ValueProps}
          eyebrow="Two sides"
          title="One pool, two things you can do with it."
          description="The interest borrowers pay is what lenders earn. There is nothing else in the middle."
        />

        <div className="mt-10 grid gap-6 lg:grid-cols-2">
          {lendingPaths.map((path) => (
            <Card key={path.key} elevation={SurfaceElevation.Flat} className="flex flex-col gap-4 p-6 sm:p-7">
              <span className="grid size-10 place-items-center rounded-tile bg-accent-soft text-accent">
                <Icon name={path.icon} className="size-5" />
              </span>

              <div className="flex flex-col gap-1">
                <span className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">{path.audience}</span>
                <h3 className="text-lg font-semibold tracking-tight text-ink">{path.title}</h3>
              </div>

              <p className="text-sm leading-relaxed text-ink-soft">{path.description}</p>

              <ol className="flex flex-col gap-2.5 border-t border-line pt-4">
                {path.steps.map((step, index) => (
                  <li key={step} className="flex items-start gap-3 text-sm leading-relaxed text-ink-soft">
                    <span className="grid size-5 shrink-0 place-items-center rounded-pill bg-surface-muted text-xs font-semibold text-ink">
                      {index + 1}
                    </span>
                    {step}
                  </li>
                ))}
              </ol>
            </Card>
          ))}
        </div>
      </Section>

      <Section id={SectionId.Trust} tone={SectionTone.Canvas}>
        <SectionHeading
          sectionId={SectionId.Trust}
          eyebrow={lendingRiskContent.eyebrow}
          title={lendingRiskContent.title}
          description={lendingRiskContent.description}
        />

        <ul className="mt-10 grid gap-5 md:grid-cols-3">
          {lendingRiskContent.points.map((point) => (
            <li key={point.title}>
              <Card elevation={SurfaceElevation.Flat} className="flex h-full flex-col gap-3 p-6">
                <span className="grid size-9 place-items-center rounded-tile bg-amber-soft text-amber">
                  <Icon name={IconName.Warning} className="size-4.5" />
                </span>
                <h3 className="text-base font-semibold tracking-tight text-ink">{point.title}</h3>
                <p className="text-sm leading-relaxed text-ink-soft">{point.body}</p>
              </Card>
            </li>
          ))}
        </ul>
      </Section>

      <Section id={SectionId.Fees} tone={SectionTone.Surface} bordered>
        <SectionHeading
          sectionId={SectionId.Fees}
          eyebrow="Costs"
          title="What it costs to use."
          description="Published in full, because a fee you find out about later is not a fee you agreed to."
        />

        <Card elevation={SurfaceElevation.Flat} className="mt-8 px-5">
          <dl className="divide-y divide-line">
            {feeItems.map((fee) => (
              <MetricRow key={fee.kind} label={fee.label} value={fee.value} hint={fee.description} emphasised />
            ))}
          </dl>
        </Card>

        <div className="mt-6">
          <Button href={AppRoute.Fusd} variant={ButtonVariant.Secondary} trailingIcon={IconName.ArrowRight}>
            Compare with minting FUSD
          </Button>
        </div>
      </Section>

      <ClosingCta />
    </>
  );
}
