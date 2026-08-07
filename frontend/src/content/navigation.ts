import { AppRoute, FooterGroupKey, NavLinkKey, NavLinkKind, SectionId } from "@/lib/enums";

export type NavLink = {
  key: NavLinkKey;
  label: string;
  kind: NavLinkKind;
  href: string;
};

export const marketingNavLinks: NavLink[] = [
  {
    key: NavLinkKey.HowItWorks,
    label: "How it works",
    kind: NavLinkKind.Anchor,
    href: `#${SectionId.HowItWorks}`,
  },
  {
    key: NavLinkKey.Fees,
    label: "Fees",
    kind: NavLinkKind.Anchor,
    href: `#${SectionId.Fees}`,
  },
  {
    key: NavLinkKey.Trust,
    label: "Security",
    kind: NavLinkKind.Anchor,
    href: `#${SectionId.Trust}`,
  },
  {
    key: NavLinkKey.Learn,
    label: "Docs",
    kind: NavLinkKind.Route,
    href: AppRoute.Learn,
  },
];

export type FooterLink = {
  label: string;
  href: AppRoute;
};

export type FooterGroup = {
  key: FooterGroupKey;
  title: string;
  links: FooterLink[];
};

export const footerGroups: FooterGroup[] = [
  {
    key: FooterGroupKey.Product,
    title: "Product",
    links: [
      { label: "Markets", href: AppRoute.Markets },
      { label: "Lend", href: AppRoute.Lend },
      { label: "Borrow", href: AppRoute.Borrow },
      { label: "Liquidations", href: AppRoute.Liquidations },
    ],
  },
  {
    key: FooterGroupKey.Learn,
    title: "Learn",
    links: [
      { label: "How it works", href: AppRoute.LearnHowItWorks },
      { label: "Health score", href: AppRoute.LearnHealthScore },
      { label: "Liquidation rules", href: AppRoute.LearnLiquidation },
      { label: "Glossary", href: AppRoute.LearnGlossary },
    ],
  },
  {
    key: FooterGroupKey.Protocol,
    title: "Protocol",
    links: [
      { label: "Fee disclosure", href: AppRoute.LearnFees },
      { label: "FAQ", href: AppRoute.LearnFaq },
      { label: "Practice mode", href: AppRoute.Practice },
      { label: "Get started", href: AppRoute.Welcome },
    ],
  },
];
