import { AppRoute, ButtonSize, ButtonVariant, IconName, SectionId } from "@/lib/enums";
import { heroContent, products } from "@/content/landing";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";

export function Hero() {
  return (
    <section
      id={SectionId.Hero}
      aria-labelledby={`${SectionId.Hero}-heading`}
      className="relative overflow-hidden border-b border-line bg-gradient-to-b from-accent-soft/70 via-canvas to-canvas"
    >
      <div
        aria-hidden="true"
        className="pointer-events-none absolute -top-40 left-1/2 size-[44rem] -translate-x-1/2 rounded-full bg-accent-muted/40 blur-3xl"
      />

      <Container className="relative flex flex-col items-center gap-7 py-20 text-center sm:py-24 lg:py-28">
        <Badge>{heroContent.eyebrow}</Badge>

        <h1
          id={`${SectionId.Hero}-heading`}
          className="max-w-4xl text-balance text-4xl font-semibold leading-[1.1] tracking-tight text-ink sm:text-5xl lg:text-6xl"
        >
          {heroContent.title}
        </h1>

        <p className="max-w-2xl text-pretty text-lg leading-relaxed text-ink-soft">{heroContent.description}</p>

        <div className="flex flex-col gap-3 sm:flex-row">
          {products.map((product) => (
            <Button
              key={product.key}
              href={product.href}
              size={ButtonSize.Lg}
              variant={product.isLive ? ButtonVariant.Primary : ButtonVariant.Secondary}
              trailingIcon={IconName.ArrowRight}
            >
              {product.title}
            </Button>
          ))}
        </div>

        <p className="flex items-center gap-2 text-sm text-ink-soft">
          <Icon name={IconName.Lock} className="size-4 text-mint" />
          {heroContent.note}
        </p>

        <Button href={AppRoute.Login} size={ButtonSize.Sm} variant={ButtonVariant.Ghost}>
          Already have a wallet? Connect it
        </Button>
      </Container>
    </section>
  );
}
