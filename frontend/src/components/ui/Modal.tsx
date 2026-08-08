"use client";

import { useEffect, type ReactNode } from "react";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";
import { Icon } from "@/components/ui/Icon";

type ModalProps = {
  open: boolean;
  onClose: () => void;
  title: string;
  children: ReactNode;
  footer?: ReactNode;
};

export function Modal({ open, onClose, title, children, footer }: ModalProps) {
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
    <div className="fixed inset-0 z-50 flex items-end justify-center p-0 sm:items-center sm:p-6">
      <button type="button" aria-label="Close dialog" onClick={onClose} className="absolute inset-0 bg-ink/25" />

      <div
        role="dialog"
        aria-modal="true"
        aria-label={title}
        className="relative flex max-h-[90dvh] w-full max-w-lg flex-col overflow-hidden rounded-t-panel border border-line bg-canvas shadow-lift sm:rounded-panel"
      >
        <div className="flex items-center justify-between gap-4 border-b border-line bg-surface px-5 py-4">
          <h2 className="text-base font-semibold tracking-tight text-ink">{title}</h2>
          <Button variant={ButtonVariant.Ghost} size={ButtonSize.Sm} onClick={onClose} ariaLabel="Close dialog">
            <Icon name={IconName.Close} className="size-4" />
          </Button>
        </div>

        <div className="flex-1 overflow-y-auto p-5">{children}</div>

        {footer ? <div className="border-t border-line bg-surface px-5 py-4">{footer}</div> : null}
      </div>
    </div>
  );
}
