import { AppNavLinkKey, ProductKey, AppRoute, FooterGroupKey, NavLinkKey, NavLinkKind } from "@/lib/enums";

export type NavLink = {
  key: NavLinkKey;
  label: string;
  kind: NavLinkKind;
  href: string;
};

export const marketingNavLinks: NavLink[] = [
  {
    key: NavLinkKey.Lending,
    label: "Lending",
    kind: NavLinkKind.Route,
    href: AppRoute.Lending,
  },
  {
    key: NavLinkKey.Fusd,
    label: "FUSD",
    kind: NavLinkKind.Route,
    href: AppRoute.Fusd,
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
    title: "Products",
    links: [
      { label: "Lending", href: AppRoute.Lending },
      { label: "FUSD", href: AppRoute.Fusd },
    ],
  },
  {
    key: FooterGroupKey.Learn,
    title: "Get started",
    links: [{ label: "Connect a wallet", href: AppRoute.Login }],
  },
];
