import { AppRoute } from "@/lib/enums";
import { Container } from "@/components/ui/Container";
import { TextLink } from "@/components/ui/TextLink";

const footerLinks = [
  { label: "Fee disclosure", href: AppRoute.LearnFees },
  { label: "How liquidation works", href: AppRoute.LearnLiquidation },
  { label: "Glossary", href: AppRoute.LearnGlossary },
  { label: "FAQ", href: AppRoute.LearnFaq },
];

export function AppFooter() {
  return (
    <footer className="border-t border-line bg-surface">
      <Container className="flex flex-col gap-4 py-8 sm:flex-row sm:items-center sm:justify-between">
        <nav aria-label="Footer" className="flex flex-wrap items-center gap-x-6 gap-y-2">
          {footerLinks.map((link) => (
            <TextLink key={link.href} href={link.href}>
              {link.label}
            </TextLink>
          ))}
        </nav>
        <p className="text-xs text-ink-faint">Reads the market straight from the chain. Nobody has audited these contracts, so do not put real money in.</p>
      </Container>
    </footer>
  );
}
