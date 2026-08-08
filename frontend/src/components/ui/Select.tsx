"use client";

import { cn } from "@/lib/cn";

export type SelectOption<TValue extends string> = {
  value: TValue;
  label: string;
};

type SelectProps<TValue extends string> = {
  id: string;
  label: string;
  value: TValue;
  options: SelectOption<TValue>[];
  onChange: (value: TValue) => void;
  className?: string;
};

export function Select<TValue extends string>({
  id,
  label,
  value,
  options,
  onChange,
  className,
}: SelectProps<TValue>) {
  return (
    <div className={cn("flex flex-col gap-1.5", className)}>
      <label htmlFor={id} className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
        {label}
      </label>
      <select
        id={id}
        value={value}
        onChange={(event) => onChange(event.target.value as TValue)}
        className="h-10 rounded-tile border border-line-strong bg-surface px-3 text-sm text-ink outline-none transition-colors hover:border-brand-border focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
      >
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
}
