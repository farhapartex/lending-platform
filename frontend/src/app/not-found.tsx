import type { Metadata } from "next";
import Link from "next/link";
import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { notFoundContent, systemDestinations } from "@/content/system";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { Logo } from "@/components/ui/Logo";

export const metadata: Metadata = {
  title: "Page not found",
  robots: { index: false, follow: false },
};

export default function NotFound() {
  return (
    <>
      <header className="border-b border-line bg-surface">
        <Container className="flex h-16 items-center">
          <Logo />
        </Container>
      </header>

      <main className="flex-1">
        <Container className="max-w-2xl py-20">
          <div className="flex flex-col items-start gap-6">
            <span className="rounded-pill border border-line bg-surface-muted px-3 py-1 text-xs font-medium text-ink-soft tabular-nums">
              404
            </span>

            <div className="flex flex-col gap-3">
              <h1 className="text-balance text-3xl font-semibold tracking-tight text-ink sm:text-4xl">
                {notFoundContent.title}
              </h1>
              <p className="text-pretty text-base leading-relaxed text-ink-soft">{notFoundContent.description}</p>
            </div>

            <div className="flex flex-col gap-3 sm:flex-row">
              <Button href={AppRoute.Markets} size={ButtonSize.Lg} trailingIcon={IconName.ArrowRight}>
                {notFoundContent.primaryCta}
              </Button>
              <Button href={AppRoute.Learn} size={ButtonSize.Lg} variant={ButtonVariant.Secondary}>
                {notFoundContent.secondaryCta}
              </Button>
            </div>

            <ul className="mt-4 grid w-full gap-3 sm:grid-cols-2">
              {systemDestinations.map((destination) => (
                <li key={destination.href}>
                  <Link
                    href={destination.href}
                    className="flex items-center justify-between gap-3 rounded-card border border-line bg-surface px-4 py-3 transition-colors hover:border-brand-border outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
                  >
                    <span className="flex flex-col gap-0.5">
                      <span className="text-sm font-medium text-ink">{destination.label}</span>
                      <span className="text-xs text-ink-faint">{destination.description}</span>
                    </span>
                    <Icon name={IconName.ArrowRight} className="size-4 shrink-0 text-ink-faint" />
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        </Container>
      </main>
    </>
  );
}
