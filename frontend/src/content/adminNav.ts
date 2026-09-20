import { AppRoute, IconName } from "@/lib/enums";

export type AdminNavItem = {
  key: string;
  label: string;
  href: AppRoute;
  icon: IconName;
};

export type AdminNavGroup = {
  key: string;
  label: string;
  items: AdminNavItem[];
};

export const adminNavGroups: AdminNavGroup[] = [
  {
    key: "overview",
    label: "Overview",
    items: [
      { key: "dashboard", label: "Dashboard", href: AppRoute.Dashboard, icon: IconName.Gauge },
      { key: "markets", label: "Markets", href: AppRoute.Markets, icon: IconName.TrendUp },
    ],
  },
  {
    key: "lending",
    label: "Lending",
    items: [
      { key: "lend", label: "Lend", href: AppRoute.Lend, icon: IconName.Coins },
      { key: "borrow", label: "Borrow", href: AppRoute.Borrow, icon: IconName.Wallet },
      { key: "liquidations", label: "Liquidations", href: AppRoute.Liquidations, icon: IconName.ShieldCheck },
    ],
  },
  {
    key: "fusd",
    label: "FUSD",
    items: [
      { key: "vault", label: "Vault", href: AppRoute.FusdVault, icon: IconName.Lock },
      { key: "backing", label: "Backing", href: AppRoute.FusdBacking, icon: IconName.Sparkles },
      { key: "fusdLiquidations", label: "Liquidations", href: AppRoute.FusdLiquidations, icon: IconName.ShieldCheck },
    ],
  },
  {
    key: "activity",
    label: "Activity",
    items: [{ key: "history", label: "History", href: AppRoute.History, icon: IconName.Receipt }],
  },
  {
    key: "help",
    label: "Help",
    items: [
      { key: "learn", label: "Learn", href: AppRoute.Learn, icon: IconName.Info },
      { key: "practice", label: "Practice mode", href: AppRoute.Practice, icon: IconName.Beaker },
    ],
  },
];
