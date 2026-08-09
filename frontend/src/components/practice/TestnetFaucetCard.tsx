import { ButtonSize, ButtonVariant, FaucetStatus, IconName, SurfaceElevation, WalletGatePurpose } from "@/lib/enums";
import { formatTokenAmount } from "@/lib/token";
import { faucetAssets, practicePageContent } from "@/content/practice";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";
import { WalletGate } from "@/components/app/WalletGate";
import { FaucetRateLimitNotice } from "@/components/practice/FaucetRateLimitNotice";

export function TestnetFaucetCard() {
  return (
    <div className="flex flex-col gap-4">
      <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
        <div className="flex flex-col gap-1 border-b border-line bg-surface-muted px-5 py-4">
          <h3 className="text-base font-semibold text-ink">{practicePageContent.faucetTitle}</h3>
          <p className="text-sm text-ink-soft">{practicePageContent.faucetDescription}</p>
        </div>

        <ul className="divide-y divide-line">
          {faucetAssets.map((asset) => (
            <li key={asset.symbol} className="flex flex-wrap items-center justify-between gap-3 px-5 py-4">
              <div className="flex flex-col gap-0.5">
                <span className="text-sm font-semibold text-ink tabular-nums">
                  {formatTokenAmount(asset.amount, asset.decimals, 2)} {asset.symbol}
                </span>
                {asset.status === FaucetStatus.CoolingDown && asset.cooldownHoursRemaining !== null ? (
                  <span className="flex items-center gap-1.5 text-xs text-amber-ink">
                    <Icon name={IconName.Info} className="size-3.5" />
                    Available again in {asset.cooldownHoursRemaining} hours
                  </span>
                ) : (
                  <span className="text-xs text-ink-faint">Ready to request</span>
                )}
              </div>

              <Button variant={ButtonVariant.Secondary} size={ButtonSize.Sm} disabled>
                {asset.status === FaucetStatus.CoolingDown ? "Requested" : "Request"}
              </Button>
            </li>
          ))}
        </ul>

        <p className="border-t border-line bg-surface-muted px-5 py-3 text-xs text-ink-faint">
          {practicePageContent.faucetPendingNote}
        </p>
      </Card>

      <WalletGate purpose={WalletGatePurpose.Faucet} skeletonClassName="h-24 rounded-card">
        <FaucetRateLimitNotice />
      </WalletGate>
    </div>
  );
}
