import { IconName } from "@/lib/enums";
import { formatSecondsAgo } from "@/lib/format";
import { lastRefreshedSecondsAgo, refreshIntervalSeconds } from "@/content/liquidations";
import { Icon } from "@/components/ui/Icon";

export function AutoRefreshIndicator() {
  return (
    <p className="flex items-center gap-2 text-xs text-ink-faint">
      <Icon name={IconName.Loader} className="size-3.5" />
      Updated {formatSecondsAgo(lastRefreshedSecondsAgo)}, refreshing every {refreshIntervalSeconds} seconds
    </p>
  );
}
