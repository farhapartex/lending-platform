"use client";

import { useEffect, useRef, useState } from "react";
import { IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

export type DropdownOption<TValue extends string> = {
  value: TValue;
  label: string;
};

type DropdownProps<TValue extends string> = {
  id: string;
  label: string;
  value: TValue;
  options: DropdownOption<TValue>[];
  onChange: (value: TValue) => void;
  className?: string;
};

export function Dropdown<TValue extends string>({
  id,
  label,
  value,
  options,
  onChange,
  className,
}: DropdownProps<TValue>) {
  const [isOpen, setIsOpen] = useState(false);
  const [activeIndex, setActiveIndex] = useState(() => Math.max(0, options.findIndex((option) => option.value === value)));

  const containerRef = useRef<HTMLDivElement>(null);
  const listRef = useRef<HTMLUListElement>(null);

  const labelId = `${id}-label`;
  const selectedOption = options.find((option) => option.value === value);

  useEffect(() => {
    if (!isOpen) {
      return;
    }

    const handlePointerDown = (event: MouseEvent) => {
      if (containerRef.current !== null && !containerRef.current.contains(event.target as Node)) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handlePointerDown);

    return () => document.removeEventListener("mousedown", handlePointerDown);
  }, [isOpen]);

  useEffect(() => {
    if (isOpen) {
      listRef.current?.focus();
    }
  }, [isOpen]);

  const open = () => {
    setActiveIndex(Math.max(0, options.findIndex((option) => option.value === value)));
    setIsOpen(true);
  };

  const commit = (index: number) => {
    const option = options[index];

    if (option !== undefined) {
      onChange(option.value);
    }

    setIsOpen(false);
  };

  const handleTriggerKeyDown = (event: React.KeyboardEvent<HTMLButtonElement>) => {
    if (event.key === "ArrowDown" || event.key === "ArrowUp" || event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      open();
    }
  };

  const handleListKeyDown = (event: React.KeyboardEvent<HTMLUListElement>) => {
    if (event.key === "Escape" || event.key === "Tab") {
      setIsOpen(false);
      return;
    }

    if (event.key === "ArrowDown") {
      event.preventDefault();
      setActiveIndex((current) => (current + 1) % options.length);
      return;
    }

    if (event.key === "ArrowUp") {
      event.preventDefault();
      setActiveIndex((current) => (current - 1 + options.length) % options.length);
      return;
    }

    if (event.key === "Home") {
      event.preventDefault();
      setActiveIndex(0);
      return;
    }

    if (event.key === "End") {
      event.preventDefault();
      setActiveIndex(options.length - 1);
      return;
    }

    if (event.key === "Enter" || event.key === " ") {
      event.preventDefault();
      commit(activeIndex);
    }
  };

  return (
    <div ref={containerRef} className={cn("flex flex-col gap-1.5", className)}>
      <span id={labelId} className="text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
        {label}
      </span>

      <div className="relative">
        <button
          type="button"
          id={id}
          aria-haspopup="listbox"
          aria-expanded={isOpen}
          aria-labelledby={`${labelId} ${id}`}
          onClick={() => (isOpen ? setIsOpen(false) : open())}
          onKeyDown={handleTriggerKeyDown}
          className="flex h-10 w-full items-center justify-between gap-3 rounded-tile border border-line-strong bg-surface px-3 text-sm text-ink transition-colors hover:border-brand-border outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas"
        >
          <span className="truncate">{selectedOption?.label ?? ""}</span>
          <Icon
            name={IconName.ChevronDown}
            className={cn("size-4 shrink-0 text-ink-faint transition-transform", isOpen && "rotate-180")}
          />
        </button>

        {isOpen ? (
          <ul
            ref={listRef}
            role="listbox"
            tabIndex={-1}
            aria-labelledby={labelId}
            aria-activedescendant={`${id}-option-${activeIndex}`}
            onKeyDown={handleListKeyDown}
            className="absolute left-0 right-0 top-full z-30 mt-1.5 overflow-hidden rounded-tile border border-line bg-surface py-1 shadow-lift outline-none"
          >
            {options.map((option, index) => {
              const isSelected = option.value === value;
              const isActive = index === activeIndex;

              return (
                <li
                  key={option.value}
                  id={`${id}-option-${index}`}
                  role="option"
                  aria-selected={isSelected}
                  onClick={() => commit(index)}
                  onMouseEnter={() => setActiveIndex(index)}
                  className={cn(
                    "flex cursor-pointer items-center justify-between gap-3 px-3 py-2 text-sm",
                    isActive ? "bg-brand-soft text-brand-ink" : "text-ink",
                  )}
                >
                  <span>{option.label}</span>
                  {isSelected ? <Icon name={IconName.Check} className="size-4 shrink-0 text-brand" /> : null}
                </li>
              );
            })}
          </ul>
        ) : null}
      </div>
    </div>
  );
}
