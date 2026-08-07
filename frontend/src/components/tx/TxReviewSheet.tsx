import type { ReactNode } from "react";

export type TxReviewRow = {
  label: string;
  value: string;
  emphasised?: boolean;
};

type TxReviewSheetProps = {
  title: string;
  rows: TxReviewRow[];
  footer?: ReactNode;
};

export function TxReviewSheet({ title, rows, footer }: TxReviewSheetProps) {
  return (
    <div className="flex flex-col gap-4 rounded-card border border-line bg-surface-muted p-5">
      <h3 className="text-sm font-semibold text-ink">{title}</h3>

      <dl className="flex flex-col gap-2.5">
        {rows.map((row) => (
          <div key={row.label} className="flex flex-wrap items-baseline justify-between gap-x-4 gap-y-0.5">
            <dt className="text-sm text-ink-soft">{row.label}</dt>
            <dd
              className={
                row.emphasised
                  ? "text-sm font-semibold text-brand-ink tabular-nums"
                  : "text-sm font-medium text-ink tabular-nums"
              }
            >
              {row.value}
            </dd>
          </div>
        ))}
      </dl>

      {footer}
    </div>
  );
}
