"use client";

import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";

type PaginationProps = {
  page: number;
  pageCount: number;
  totalItems: number;
  pageSize: number;
  onPageChange: (page: number) => void;
};

export function Pagination({ page, pageCount, totalItems, pageSize, onPageChange }: PaginationProps) {
  const firstItem = totalItems === 0 ? 0 : (page - 1) * pageSize + 1;
  const lastItem = Math.min(page * pageSize, totalItems);

  return (
    <nav
      aria-label="Pagination"
      className="flex flex-wrap items-center justify-between gap-3 border-t border-line bg-surface-muted px-5 py-3.5"
    >
      <p className="text-xs text-ink-soft tabular-nums">
        Showing {firstItem} to {lastItem} of {totalItems}
      </p>

      <div className="flex items-center gap-2">
        <Button
          variant={ButtonVariant.Secondary}
          size={ButtonSize.Sm}
          disabled={page <= 1}
          onClick={() => onPageChange(page - 1)}
          ariaLabel="Previous page"
        >
          <Icon name={IconName.ArrowRight} className="size-4 rotate-180" />
        </Button>

        <span className="text-xs text-ink-soft tabular-nums">
          Page {page} of {pageCount}
        </span>

        <Button
          variant={ButtonVariant.Secondary}
          size={ButtonSize.Sm}
          disabled={page >= pageCount}
          onClick={() => onPageChange(page + 1)}
          ariaLabel="Next page"
        >
          <Icon name={IconName.ArrowRight} className="size-4" />
        </Button>
      </div>
    </nav>
  );
}
