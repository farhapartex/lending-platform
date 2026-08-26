"use client";

import { AppRoute, BadgeTone, ButtonSize, ButtonVariant, IconName, SurfaceElevation } from "@/lib/enums";
import { formatDateTimeUtc } from "@/lib/format";
import { formatTokenAmount } from "@/lib/token";
import type { TransactionDetail } from "@/lib/api/transactionMapper";
import { dashboardContent } from "@/content/dashboard";
import { useRecentActivity } from "@/hooks/useRecentActivity";
import { useWalletState } from "@/hooks/useWalletState";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { TextLink } from "@/components/ui/TextLink";
import { TxTypeBadge } from "@/components/tx/TxTypeBadge";

export function RecentActivityList() {
  const { address } = useWalletState();
  const { page, isLoading, isError, refetch } = useRecentActivity(address);

  if (isLoading) {
    return <ActivitySkeleton />;
  }

  if (isError) {
    return (
      <Alert title={dashboardContent.activityUnavailableTitle} tone={BadgeTone.Neutral} icon={IconName.Info}>
        <div className="flex flex-col items-start gap-2">
          <p>{dashboardContent.activityUnavailableDescription}</p>
          <Button variant={ButtonVariant.Subtle} size={ButtonSize.Sm} onClick={refetch}>
            {dashboardContent.activityRetry}
          </Button>
        </div>
      </Alert>
    );
  }

  const entries = page?.items ?? [];

  if (entries.length === 0) {
    return <ActivityEmptyState indexed={page?.asOf.block !== null} />;
  }

  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
      <ul className="divide-y divide-line">
        {entries.map((entry) => (
          <ActivityRow key={entry.id} entry={entry} />
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

function ActivityRow({ entry }: { entry: TransactionDetail }) {
  return (
    <li className="flex flex-wrap items-center justify-between gap-3 px-5 py-4 sm:px-6">
      <div className="flex flex-col gap-1.5">
        <TxTypeBadge kind={entry.kind} />
        <span className="text-xs text-ink-faint">{formatDateTimeUtc(entry.timestamp)}</span>
      </div>
      <span className="text-sm font-semibold text-ink tabular-nums">
        {formatTokenAmount(entry.amount, entry.decimals, 4)} {entry.symbol}
      </span>
    </li>
  );
}

function ActivityEmptyState({ indexed }: { indexed: boolean }) {
  if (!indexed) {
    return (
      <EmptyState
        title={dashboardContent.activityNotIndexedTitle}
        description={dashboardContent.activityNotIndexedDescription}
        icon={IconName.Info}
      />
    );
  }

  return (
    <EmptyState
      title={dashboardContent.activityEmptyTitle}
      description={dashboardContent.activityEmptyDescription}
      icon={IconName.Receipt}
    />
  );
}

function ActivitySkeleton() {
  return (
    <Card elevation={SurfaceElevation.Raised} className="overflow-hidden">
      <ul className="divide-y divide-line">
        {[0, 1, 2].map((row) => (
          <li key={row} className="flex items-center justify-between gap-3 px-5 py-4 sm:px-6">
            <div className="flex flex-col gap-2">
              <Skeleton className="h-5 w-24" />
              <Skeleton className="h-3 w-32" />
            </div>
            <Skeleton className="h-4 w-20" />
          </li>
        ))}
      </ul>
    </Card>
  );
}
