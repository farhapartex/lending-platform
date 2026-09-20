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
import {
  fusdPageContent,
  fusdSteps,
  fusdTerms,
  fusdUses,
  healthContent,
  howItWorksContent,
  limitsContent,
  pegContent,
  stablecoinKinds,
  statusContent,
  usesContent,
  whatItIsContent,
} from "@/content/fusd";
import { Alert } from "@/components/ui/Alert";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { MetricRow } from "@/components/ui/MetricRow";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

export const metadata: Metadata = {
  title: "FUSD",
  description:
    "FUSD is a decentralized, overcollateralized, dollar-pegged stablecoin. Lock crypto you already own and mint dollars against it, with no company holding reserves and no freeze switch.",
};

export default function FusdPage() {
  return (
    <>
      <section
        id={SectionId.Hero}
        aria-labelledby={`${SectionId.Hero}-heading`}
        className="relative overflow-hidden border-b border-line bg-gradient-to-b from-brand-soft/70 via-canvas to-canvas"
      >
        <div
          aria-hidden="true"
          className="pointer-events-none absolute -top-32 left-1/2 size-[42rem] -translate-x-1/2 rounded-full bg-brand-muted/45 blur-3xl"
        />
        <Container className="relative flex flex-col items-start gap-6 py-16 sm:py-20 lg:py-24">
          <Badge tone={BadgeTone.Caution}>{fusdPageContent.eyebrow}</Badge>

          <h1
            id={`${SectionId.Hero}-heading`}
            className="max-w-3xl text-balance text-4xl font-semibold leading-[1.1] tracking-tight text-ink sm:text-5xl"
          >
            {fusdPageContent.title}
          </h1>

          <p className="max-w-2xl text-pretty text-lg leading-relaxed text-ink-soft">
            {fusdPageContent.description}
          </p>

          <Alert title="Not live yet" tone={BadgeTone.Caution} icon={IconName.Info}>
            {fusdPageContent.statusNote}
          </Alert>
        </Container>
      </section>

      <Section id={SectionId.FusdWhat} tone={SectionTone.Canvas}>
        <SectionHeading
          sectionId={SectionId.FusdWhat}
          eyebrow={whatItIsContent.eyebrow}
          title={whatItIsContent.title}
          description={whatItIsContent.description}
        />

        <div className="mt-10 grid gap-6 lg:grid-cols-3">
          {stablecoinKinds.map((kind) => {
            const isChosen = kind.key === "crypto";

            return (
              <Card
                key={kind.key}
                elevation={isChosen ? SurfaceElevation.Raised : SurfaceElevation.Flat}
                className="flex flex-col gap-3 p-6"
              >
                <Badge tone={isChosen ? BadgeTone.Positive : BadgeTone.Neutral}>{kind.verdict}</Badge>
                <h3 className="text-base font-semibold tracking-tight text-ink">{kind.title}</h3>
                <p className="text-sm leading-relaxed text-ink-soft">{kind.description}</p>
              </Card>
            );
          })}
        </div>
      </Section>

      <Section id={SectionId.FusdHow} tone={SectionTone.Surface} bordered>
        <SectionHeading
          sectionId={SectionId.FusdHow}
          eyebrow={howItWorksContent.eyebrow}
          title={howItWorksContent.title}
          description={howItWorksContent.description}
        />

        <ol className="mt-10 grid gap-6 sm:grid-cols-2">
          {fusdSteps.map((step, index) => (
            <li key={step.key} className="flex gap-4 rounded-card border border-line bg-surface p-6">
              <span className="flex size-8 shrink-0 items-center justify-center rounded-pill bg-brand-soft text-sm font-semibold text-brand-ink">
                {index + 1}
              </span>
              <div className="flex flex-col gap-2">
                <h3 className="text-base font-semibold tracking-tight text-ink">{step.title}</h3>
                <p className="text-sm leading-relaxed text-ink-soft">{step.description}</p>
              </div>
            </li>
          ))}
        </ol>
      </Section>

      <Section id={SectionId.FusdTerms} tone={SectionTone.Canvas}>
        <SectionHeading
          sectionId={SectionId.FusdTerms}
          eyebrow={healthContent.eyebrow}
          title={healthContent.title}
          description={healthContent.description}
        />

        <div className="mt-10 grid gap-6 lg:grid-cols-2 lg:gap-8">
          <div className="flex flex-col gap-5">
            <Card elevation={SurfaceElevation.Flat} className="p-6">
              <code className="block text-sm leading-relaxed text-brand-ink">{healthContent.formula}</code>
            </Card>

            <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-2 p-6">
              <h3 className="text-base font-semibold tracking-tight text-ink">{healthContent.exampleTitle}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{healthContent.exampleBody}</p>
            </Card>

            <p className="flex items-start gap-2 text-sm text-ink-soft">
              <Icon name={IconName.Info} className="mt-0.5 size-4 text-mint" />
              {healthContent.drift}
            </p>
          </div>

          <Card elevation={SurfaceElevation.Raised} className="p-6 sm:p-7">
            <h3 className="text-base font-semibold tracking-tight text-ink">Published terms</h3>
            <dl className="mt-4 divide-y divide-line">
              {fusdTerms.map((term) => (
                <MetricRow key={term.label} label={term.label} value={term.value} hint={term.hint} />
              ))}
            </dl>
          </Card>
        </div>
      </Section>

      <Section id={SectionId.FusdUses} tone={SectionTone.Surface} bordered>
        <SectionHeading sectionId={SectionId.FusdUses} eyebrow={usesContent.eyebrow} title={usesContent.title} />

        <div className="mt-10 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          {fusdUses.map((use) => (
            <Card key={use.key} elevation={SurfaceElevation.Flat} className="flex flex-col gap-3 p-6">
              <span className="flex size-10 items-center justify-center rounded-tile bg-brand-soft">
                <Icon name={use.icon} className="size-5 text-brand-ink" />
              </span>
              <h3 className="text-base font-semibold tracking-tight text-ink">{use.title}</h3>
              <p className="text-sm leading-relaxed text-ink-soft">{use.description}</p>
            </Card>
          ))}
        </div>
      </Section>

      <Section id={SectionId.FusdPeg} tone={SectionTone.Canvas}>
        <SectionHeading sectionId={SectionId.FusdPeg} eyebrow={pegContent.eyebrow} title={pegContent.title} />

        <div className="mt-10 grid gap-6 lg:grid-cols-3">
          <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-3 p-6">
            <h3 className="text-base font-semibold tracking-tight text-ink">Above a dollar</h3>
            <p className="text-sm leading-relaxed text-ink-soft">{pegContent.above}</p>
          </Card>
          <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-3 p-6">
            <h3 className="text-base font-semibold tracking-tight text-ink">Below a dollar</h3>
            <p className="text-sm leading-relaxed text-ink-soft">{pegContent.below}</p>
          </Card>
          <Card elevation={SurfaceElevation.Raised} className="flex flex-col gap-3 p-6">
            <h3 className="text-base font-semibold tracking-tight text-ink">Underneath both</h3>
            <p className="text-sm leading-relaxed text-ink-soft">{pegContent.floor}</p>
          </Card>
        </div>

        <Card elevation={SurfaceElevation.Flat} className="mt-8 flex flex-col gap-4 p-6 sm:p-7">
          <div className="flex flex-col gap-1.5">
            <span className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
              {limitsContent.eyebrow}
            </span>
            <h3 className="text-lg font-semibold tracking-tight text-ink">{limitsContent.title}</h3>
          </div>
          <ul className="flex flex-col gap-2">
            {limitsContent.items.map((item) => (
              <li key={item} className="flex items-start gap-2 text-sm leading-relaxed text-ink-soft">
                <Icon name={IconName.Warning} className="mt-0.5 size-4 shrink-0 text-amber" />
                {item}
              </li>
            ))}
          </ul>
        </Card>
      </Section>

      <Section id={SectionId.FusdStatus} tone={SectionTone.Surface} bordered>
        <SectionHeading
          sectionId={SectionId.FusdStatus}
          eyebrow={statusContent.eyebrow}
          title={statusContent.title}
        />

        <div className="mt-10 grid gap-6 lg:grid-cols-3">
          {[
            { heading: "Built", items: statusContent.done, tone: BadgeTone.Positive, label: "Done" },
            { heading: "Being written", items: statusContent.next, tone: BadgeTone.Caution, label: "Next" },
            { heading: "After that", items: statusContent.after, tone: BadgeTone.Neutral, label: "Later" },
          ].map((group) => (
            <Card key={group.heading} elevation={SurfaceElevation.Flat} className="flex flex-col gap-4 p-6">
              <div className="flex items-center justify-between gap-3">
                <h3 className="text-base font-semibold tracking-tight text-ink">{group.heading}</h3>
                <Badge tone={group.tone}>{group.label}</Badge>
              </div>
              <ul className="flex flex-col gap-2.5">
                {group.items.map((item) => (
                  <li key={item} className="text-sm leading-relaxed text-ink-soft">
                    {item}
                  </li>
                ))}
              </ul>
            </Card>
          ))}
        </div>

        <div className="mt-10 flex flex-col gap-3 sm:flex-row">
          <Button href={AppRoute.Lending} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight}>
            See what is live today
          </Button>
          <Button href={AppRoute.Login} size={ButtonSize.Lg} variant={ButtonVariant.Secondary}>
            Connect a wallet
          </Button>
        </div>
      </Section>
    </>
  );
}
