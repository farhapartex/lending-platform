"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { AppRoute } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { docPages } from "@/content/learn";

const baseClasses =
  "block rounded-tile px-3 py-2 text-sm transition-colors outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

export function DocsSidebar() {
  const pathname = usePathname();

  return (
    <nav aria-label="Documentation" className="flex flex-col gap-1">
      <Link
        href={AppRoute.Learn}
        aria-current={pathname === AppRoute.Learn ? "page" : undefined}
        className={cn(
          baseClasses,
          pathname === AppRoute.Learn ? "bg-brand-soft font-medium text-brand-ink" : "text-ink-soft hover:bg-surface-muted hover:text-ink",
        )}
      >
        Overview
      </Link>

      {docPages.map((page) => {
        const isActive = pathname === page.route;

        return (
          <Link
            key={page.key}
            href={page.route}
            aria-current={isActive ? "page" : undefined}
            className={cn(
              baseClasses,
              isActive ? "bg-brand-soft font-medium text-brand-ink" : "text-ink-soft hover:bg-surface-muted hover:text-ink",
            )}
          >
            {page.title}
          </Link>
        );
      })}
    </nav>
  );
}
