import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Button } from "@/components/ui/Button";

type MarketActionsBarProps = {
  className?: string;
};

export function MarketActionsBar({ className }: MarketActionsBarProps) {
  return (
    <div className={cn("flex flex-col gap-3 sm:flex-row", className)}>
      <Button href={AppRoute.Lend} size={ButtonSize.Md} trailingIcon={IconName.ArrowRight}>
        Lend USDC
      </Button>
      <Button href={AppRoute.Borrow} size={ButtonSize.Md} variant={ButtonVariant.Secondary}>
        Borrow against WETH
      </Button>
    </div>
  );
}
