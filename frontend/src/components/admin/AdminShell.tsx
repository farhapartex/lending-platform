"use client";

import { useState, type ReactNode } from "react";
import { IconName, SectionId } from "@/lib/enums";
import { AdminSidebar } from "@/components/admin/AdminSidebar";
import { AdminTopBar } from "@/components/admin/AdminTopBar";
import { Icon } from "@/components/ui/Icon";

const drawerId = "admin-sidebar-drawer";

export function AdminShell({ children }: { children: ReactNode }) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="flex min-h-screen bg-canvas">
      <aside className="hidden w-64 shrink-0 lg:block">
        <div className="fixed inset-y-0 left-0 w-64">
          <AdminSidebar />
        </div>
      </aside>

      {isOpen ? (
        <div className="fixed inset-0 z-50 lg:hidden">
          <button
            type="button"
            aria-label="Close the menu"
            onClick={() => setIsOpen(false)}
            className="absolute inset-0 bg-ink/30"
          />
          <div id={drawerId} className="absolute inset-y-0 left-0 w-64">
            <AdminSidebar onNavigate={() => setIsOpen(false)} />
          </div>
          <button
            type="button"
            aria-label="Close the menu"
            onClick={() => setIsOpen(false)}
            className="absolute right-4 top-4 grid size-9 place-items-center rounded-pill bg-surface text-ink shadow-soft"
          >
            <Icon name={IconName.Close} className="size-4" />
          </button>
        </div>
      ) : null}

      <div className="flex min-w-0 flex-1 flex-col">
        <AdminTopBar onOpenSidebar={() => setIsOpen(true)} />
        <main id={SectionId.MainContent} className="flex-1">
          {children}
        </main>
      </div>
    </div>
  );
}
