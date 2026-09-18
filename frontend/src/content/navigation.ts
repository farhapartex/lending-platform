import { AppNavLinkKey, ProductKey, AppRoute, FooterGroupKey, NavLinkKey, NavLinkKind, SectionId } from "@/lib/enums";

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
    key: NavLinkKey.BorrowOrMint,
    label: "Ways to borrow",
    kind: NavLinkKind.Anchor,
    href: `#${SectionId.BorrowOrMint}`,
  },
  {
    key: NavLinkKey.Fusd,
    label: "FUSD",
    kind: NavLinkKind.Route,
    href: AppRoute.Fusd,
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

export type AppNavLink = {
  key: AppNavLinkKey;
  label: string;
  href: AppRoute;
};

export type NavProduct = {
  key: ProductKey;
  label: string;
  home: AppRoute;
  links: AppNavLink[];
};

const sharedLinks: AppNavLink[] = [
  { key: AppNavLinkKey.Dashboard, label: "Dashboard", href: AppRoute.Dashboard },
  { key: AppNavLinkKey.History, label: "History", href: AppRoute.History },
  { key: AppNavLinkKey.Learn, label: "Learn", href: AppRoute.Learn },
];

export const navProducts: NavProduct[] = [
  {
    key: ProductKey.Lending,
    label: "Lending",
    home: AppRoute.Markets,
    links: [
      { key: AppNavLinkKey.Markets, label: "Markets", href: AppRoute.Markets },
      { key: AppNavLinkKey.Lend, label: "Lend", href: AppRoute.Lend },
      { key: AppNavLinkKey.Borrow, label: "Borrow", href: AppRoute.Borrow },
      { key: AppNavLinkKey.Liquidations, label: "Liquidations", href: AppRoute.Liquidations },
      ...sharedLinks,
    ],
  },
  {
    key: ProductKey.Fusd,
    label: "FUSD",
    home: AppRoute.FusdVault,
    links: [
      { key: AppNavLinkKey.FusdVault, label: "Vault", href: AppRoute.FusdVault },
      { key: AppNavLinkKey.FusdLiquidations, label: "Liquidations", href: AppRoute.FusdLiquidations },
      { key: AppNavLinkKey.FusdBacking, label: "Backing", href: AppRoute.FusdBacking },
      ...sharedLinks,
    ],
  },
];

export function productForPath(pathname: string): NavProduct {
  const fusd = navProducts[1];

  return pathname.startsWith("/fusd") ? fusd : navProducts[0];
}

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
