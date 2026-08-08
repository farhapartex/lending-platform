import { AppRoute, IconName, SurfaceElevation } from "@/lib/enums";
import { formatDateTimeUtc } from "@/lib/format";
import { formatTokenAmount } from "@/lib/token";
import { recentActivity } from "@/content/dashboard";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { TextLink } from "@/components/ui/TextLink";
import { TxTypeBadge } from "@/components/tx/TxTypeBadge";

export function RecentActivityList() {
  if (recentActivity.length === 0) {
    return (
      <EmptyState
        title="Nothing has happened yet"
        description="Your deposits, loans, repayments, and any liquidations will show up here."
        icon={IconName.Receipt}
      />
    );
  }

  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
      <ul className="divide-y divide-line">
        {recentActivity.map((entry) => (
          <li key={entry.id} className="flex flex-wrap items-center justify-between gap-3 px-5 py-4 sm:px-6">
            <div className="flex flex-col gap-1.5">
              <TxTypeBadge kind={entry.kind} />
              <span className="text-xs text-ink-faint">{formatDateTimeUtc(entry.timestamp)}</span>
            </div>
            <span className="text-sm font-semibold text-ink tabular-nums">
              {formatTokenAmount(entry.amount, entry.decimals, 4)} {entry.symbol}
            </span>
          </li>
        ))}
      </ul>

      <div className="border-t border-line bg-surface-muted px-5 py-3.5 sm:px-6">
        <TextLink href={AppRoute.History} trailingIcon={IconName.ArrowRight}>
          See full transaction history
        </TextLink>
      </div>
    </Card>
  );
}
