import { footerGroups } from "@/content/navigation";
import { Container } from "@/components/ui/Container";
import { Logo } from "@/components/ui/Logo";
import { TextLink } from "@/components/ui/TextLink";

export function MarketingFooter() {
  return (
    <footer className="border-t border-line bg-surface">
      <Container className="py-14">
        <div className="grid gap-10 lg:grid-cols-[minmax(0,20rem)_minmax(0,1fr)] lg:gap-16">
          <div className="flex flex-col gap-4">
            <Logo />
            <p className="max-w-xs text-sm leading-relaxed text-ink-soft">
              A transparent, non-custodial market for lending and borrowing crypto. You keep custody of your assets at
              every step.
            </p>
          </div>

          <nav aria-label="Footer" className="grid gap-8 sm:grid-cols-3">
            {footerGroups.map((group) => (
              <div key={group.key} className="flex flex-col gap-3">
                <h2 className="text-xs font-semibold uppercase tracking-[0.12em] text-ink-faint">{group.title}</h2>
                <ul className="flex flex-col gap-2.5">
                  {group.links.map((link) => (
                    <li key={link.href}>
                      <TextLink href={link.href}>{link.label}</TextLink>
                    </li>
                  ))}
                </ul>
              </div>
            ))}
          </nav>
        </div>

        <div className="mt-12 flex flex-col gap-4 border-t border-line pt-8 sm:flex-row sm:items-center sm:justify-between">
          <p className="text-xs text-ink-faint">Interface preview built with mock data. Not yet connected to a live market.</p>
          <p className="text-xs text-ink-faint">
            Lending and borrowing crypto carries risk, including liquidation of your collateral.
          </p>
        </div>
      </Container>
    </footer>
  );
}
