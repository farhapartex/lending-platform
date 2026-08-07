"use client";

import { cn } from "@/lib/cn";

export type TabItem<TValue extends string> = {
  value: TValue;
  label: string;
};

type TabBarProps<TValue extends string> = {
  items: TabItem<TValue>[];
  active: TValue;
  onChange: (value: TValue) => void;
  label: string;
  className?: string;
};

export function TabBar<TValue extends string>({ items, active, onChange, label, className }: TabBarProps<TValue>) {
  return (
    <div
      role="tablist"
      aria-label={label}
      className={cn("inline-flex rounded-pill border border-line bg-surface-muted p-1", className)}
    >
      {items.map((item) => {
        const isActive = item.value === active;

        return (
          <button
            key={item.value}
            type="button"
            role="tab"
            aria-selected={isActive}
            onClick={() => onChange(item.value)}
            className={cn(
              "rounded-pill px-4 py-2 text-sm font-medium transition-colors outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-surface",
              isActive ? "bg-surface text-ink shadow-soft" : "text-ink-soft hover:text-ink",
            )}
          >
            {item.label}
          </button>
        );
      })}
    </div>
  );
}
