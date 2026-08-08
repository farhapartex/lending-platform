"use client";

import { useEffect, type ReactNode } from "react";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";

type DrawerProps = {
  open: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
};

export function Drawer({ open, onClose, title, children }: DrawerProps) {
  useEffect(() => {
    if (!open) {
      return;
    }

    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        onClose();
      }
    };

    window.addEventListener("keydown", handleKeyDown);

    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [open, onClose]);

  if (!open) {
    return null;
  }

  return (
    <div className="fixed inset-0 z-50 flex justify-end">
      <button type="button" aria-label="Close panel" onClick={onClose} className="absolute inset-0 bg-ink/25" />

      <div
        role="dialog"
        aria-modal="true"
        aria-label={title}
        className="relative flex h-full w-full max-w-lg flex-col overflow-y-auto border-l border-line bg-canvas shadow-lift"
      >
        <div className="flex items-center justify-between gap-4 border-b border-line bg-surface px-5 py-4">
          <h2 className="text-base font-semibold tracking-tight text-ink">{title}</h2>
          <Button variant={ButtonVariant.Ghost} size={ButtonSize.Sm} onClick={onClose} ariaLabel="Close panel">
            <Icon name={IconName.Close} className="size-4" />
          </Button>
        </div>

        <div className="p-5">{children}</div>
      </div>
    </div>
  );
}
