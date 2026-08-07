"use client";

import Link from "next/link";
import { useState } from "react";
import { AppRoute, ButtonSize, ButtonVariant, IconName, NavLinkKind } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { marketingNavLinks } from "@/content/navigation";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { Logo } from "@/components/ui/Logo";

const mobilePanelId = "marketing-nav-panel";

const linkClasses =
  "rounded-pill px-3 py-2 text-sm text-ink-soft transition-colors hover:bg-surface-muted hover:text-ink outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

export function MarketingNav() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <header className="sticky top-0 z-40 border-b border-line/80 bg-canvas/85 backdrop-blur-md">
      <Container className="flex h-16 items-center justify-between gap-4">
        <Logo />

        <nav aria-label="Main" className="hidden items-center gap-1 md:flex">
          {marketingNavLinks.map((link) =>
            link.kind === NavLinkKind.Anchor ? (
              <a key={link.key} href={link.href} className={linkClasses}>
                {link.label}
              </a>
            ) : (
              <Link key={link.key} href={link.href} className={linkClasses}>
                {link.label}
              </Link>
            ),
          )}
        </nav>

        <div className="hidden md:block">
          <Button href={AppRoute.Markets} size={ButtonSize.Sm} trailingIcon={IconName.ArrowRight}>
            Launch app
          </Button>
        </div>

        <Button
          variant={ButtonVariant.Secondary}
          size={ButtonSize.Sm}
          className="md:hidden"
          onClick={() => setIsOpen(!isOpen)}
          ariaLabel={isOpen ? "Close menu" : "Open menu"}
          ariaExpanded={isOpen}
          ariaControls={mobilePanelId}
        >
          <Icon name={isOpen ? IconName.Close : IconName.Menu} className="size-4" />
        </Button>
      </Container>

      <div id={mobilePanelId} hidden={!isOpen} className={cn("border-t border-line bg-surface md:hidden")}>
        <Container className="flex flex-col gap-1 py-4">
          {marketingNavLinks.map((link) =>
            link.kind === NavLinkKind.Anchor ? (
              <a key={link.key} href={link.href} className={linkClasses} onClick={() => setIsOpen(false)}>
                {link.label}
              </a>
            ) : (
              <Link key={link.key} href={link.href} className={linkClasses} onClick={() => setIsOpen(false)}>
                {link.label}
              </Link>
            ),
          )}
          <Button href={AppRoute.Markets} fullWidth trailingIcon={IconName.ArrowRight} className="mt-3">
            Launch app
          </Button>
        </Container>
      </div>
    </header>
  );
}
