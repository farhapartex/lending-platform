import { AppRoute, BadgeTone, ButtonSize, ButtonVariant, IconName, SectionId } from "@/lib/enums";
import { heroContent } from "@/content/landing";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { MarketSnapshotCard } from "@/components/marketing/MarketSnapshotCard";

export function Hero() {
  return (
    <section
      id={SectionId.Hero}
      aria-labelledby={`${SectionId.Hero}-heading`}
      className="relative overflow-hidden border-b border-line bg-gradient-to-b from-brand-soft/70 via-canvas to-canvas"
    >
      <div
        aria-hidden="true"
        className="pointer-events-none absolute -top-32 left-1/2 size-[42rem] -translate-x-1/2 rounded-full bg-brand-muted/45 blur-3xl"
      />
      <Container className="relative grid items-center gap-12 py-16 sm:py-20 lg:grid-cols-[minmax(0,1fr)_26rem] lg:gap-16 lg:py-24">
        <div className="flex flex-col items-start gap-6">
          <Badge tone={BadgeTone.Brand}>{heroContent.eyebrow}</Badge>

          <h1
            id={`${SectionId.Hero}-heading`}
            className="text-balance text-4xl font-semibold leading-[1.1] tracking-tight text-ink sm:text-5xl"
          >
            {heroContent.title}
          </h1>

          <p className="max-w-xl text-pretty text-lg leading-relaxed text-ink-soft">{heroContent.description}</p>

          <div className="flex flex-col gap-3 sm:flex-row">
            <Button href={AppRoute.Lend} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight}>
              {heroContent.primaryCta}
            </Button>
            <Button href={AppRoute.Borrow} size={ButtonSize.Lg} variant={ButtonVariant.Secondary}>
              {heroContent.secondaryCta}
            </Button>
          </div>

          <p className="flex items-start gap-2 text-sm text-ink-soft">
            <Icon name={IconName.Lock} className="mt-0.5 size-4 text-mint" />
            {heroContent.custodyNote}
          </p>
        </div>

        <div className="w-full justify-self-end">
          <MarketSnapshotCard />
        </div>
      </Container>
    </section>
  );
}
