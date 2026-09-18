"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useState } from "react";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { productForPath } from "@/content/navigation";
import { Button } from "@/components/ui/Button";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { Logo } from "@/components/ui/Logo";
import { NetworkBadge } from "@/components/app/NetworkBadge";
import { ProductSwitcher } from "@/components/app/ProductSwitcher";
import { WalletConnectButton } from "@/components/app/WalletConnectButton";

const mobilePanelId = "app-nav-panel";

const baseLinkClasses =
  "whitespace-nowrap rounded-pill px-3 py-2 text-sm transition-colors outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

const inactiveLinkClasses = "text-ink-soft hover:bg-surface-muted hover:text-ink";

const activeLinkClasses = "bg-brand-soft font-medium text-brand-ink";

export function TopNav() {
  const [isOpen, setIsOpen] = useState(false);
  const pathname = usePathname();

  const isActive = (href: string) => pathname === href || pathname.startsWith(`${href}/`);
  const product = productForPath(pathname);

  return (
    <header className="sticky top-0 z-40 border-b border-line bg-canvas/90 backdrop-blur-md">
      <Container className="flex h-16 items-center justify-between gap-4">
        <div className="flex items-center gap-6">
          <Logo labelClassName="lg:hidden xl:inline" />
          <ProductSwitcher active={product} className="hidden lg:flex" />
          <nav aria-label="Main" className="hidden items-center gap-1 lg:flex">
            {product.links.map((link) => (
              <Link
                key={link.key}
                href={link.href}
                aria-current={isActive(link.href) ? "page" : undefined}
                className={cn(baseLinkClasses, isActive(link.href) ? activeLinkClasses : inactiveLinkClasses)}
              >
                {link.label}
              </Link>
            ))}
          </nav>
        </div>

        <div className="hidden items-center gap-3 md:flex">
          <NetworkBadge />
          <WalletConnectButton />
        </div>

        <Button
          variant={ButtonVariant.Secondary}
          size={ButtonSize.Sm}
          className="lg:hidden"
          onClick={() => setIsOpen(!isOpen)}
          ariaLabel={isOpen ? "Close menu" : "Open menu"}
          ariaExpanded={isOpen}
          ariaControls={mobilePanelId}
        >
          <Icon name={isOpen ? IconName.Close : IconName.Menu} className="size-4" />
        </Button>
      </Container>

      <div id={mobilePanelId} hidden={!isOpen} className="border-t border-line bg-surface lg:hidden">
        <Container className="flex flex-col gap-1 py-4">
          <ProductSwitcher active={product} onNavigate={() => setIsOpen(false)} className="mb-2 self-start" />
          {product.links.map((link) => (
            <Link
              key={link.key}
              href={link.href}
              aria-current={isActive(link.href) ? "page" : undefined}
              onClick={() => setIsOpen(false)}
              className={cn(baseLinkClasses, isActive(link.href) ? activeLinkClasses : inactiveLinkClasses)}
            >
              {link.label}
            </Link>
          ))}
          <div className="mt-3 flex items-center gap-3 md:hidden">
            <NetworkBadge />
            <WalletConnectButton />
          </div>
        </Container>
      </div>
    </header>
  );
}
