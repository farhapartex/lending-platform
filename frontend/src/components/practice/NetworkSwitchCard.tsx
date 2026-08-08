import { ButtonVariant, IconName, NetworkKind, SurfaceElevation } from "@/lib/enums";
import { activeNetwork } from "@/content/protocol";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";

const isOnTestnet = activeNetwork === NetworkKind.Testnet;

export function NetworkSwitchCard() {
  return (
    <Card elevation={SurfaceElevation.Flat} className="flex flex-col gap-4 p-6">
      <div className="flex items-start gap-3">
        <span
          className={
            isOnTestnet
              ? "grid size-9 shrink-0 place-items-center rounded-tile bg-mint-soft text-mint-ink"
              : "grid size-9 shrink-0 place-items-center rounded-tile bg-surface-muted text-ink-soft"
          }
        >
          <Icon name={isOnTestnet ? IconName.Check : IconName.Sliders} className="size-4.5" />
        </span>

        <div className="flex flex-col gap-1">
          <h3 className="text-base font-semibold text-ink">
            {isOnTestnet ? "You are on the test network" : "Switch to the test network"}
          </h3>
          <p className="text-sm leading-relaxed text-ink-soft">
            {isOnTestnet
              ? "Nothing you do from here touches real funds. Your live positions are untouched and waiting."
              : "Your wallet is pointed at the live network. Switch it over to practise without risking anything."}
          </p>
        </div>
      </div>

      {isOnTestnet ? null : (
        <Button variant={ButtonVariant.Secondary} disabled>
          Switch network
        </Button>
      )}
    </Card>
  );
}
