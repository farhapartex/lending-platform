import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";

const actions = [
  { label: "Deposit USDC", href: AppRoute.Lend, icon: IconName.Coins },
  { label: "Borrow USDC", href: AppRoute.Borrow, icon: IconName.Wallet },
  { label: "Repay loan", href: AppRoute.Borrow, icon: IconName.Check },
  { label: "Add collateral", href: AppRoute.Borrow, icon: IconName.ShieldCheck },
];

export function QuickActionBar() {
  return (
    <div className="flex flex-wrap gap-2">
      {actions.map((action) => (
        <Button
          key={action.label}
          href={action.href}
          size={ButtonSize.Sm}
          variant={ButtonVariant.Secondary}
          leadingIcon={action.icon}
        >
          {action.label}
        </Button>
      ))}
    </div>
  );
}
