"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { AppRoute } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { adminNavGroups } from "@/content/adminNav";
import { Icon } from "@/components/ui/Icon";
import { Logo } from "@/components/ui/Logo";

const itemClasses =
  "flex items-center gap-3 rounded-tile px-3 py-2 text-sm transition-colors outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-shell";

const inactiveClasses = "text-ink-soft hover:bg-shell-muted hover:text-ink";

const activeClasses = "bg-accent-soft font-medium text-accent-ink";

type AdminSidebarProps = {
  onNavigate?: () => void;
};

export function AdminSidebar({ onNavigate }: AdminSidebarProps) {
  const pathname = usePathname();

  const isActive = (href: AppRoute) => pathname === href || pathname.startsWith(`${href}/`);

  return (
    <div className="flex h-full flex-col gap-6 overflow-y-auto border-r border-shell-line bg-shell px-4 py-5">
      <div className="px-2">
        <Logo />
      </div>

      <nav aria-label="Sections" className="flex flex-1 flex-col gap-6">
        {adminNavGroups.map((group) => (
          <div key={group.key} className="flex flex-col gap-1">
            <span className="px-3 pb-1 text-[0.6875rem] font-semibold uppercase tracking-[0.1em] text-ink-faint">
              {group.label}
            </span>

            {group.items.map((item) => {
              const active = isActive(item.href);

              return (
                <Link
                  key={item.key}
                  href={item.href}
                  aria-current={active ? "page" : undefined}
                  onClick={onNavigate}
                  className={cn(itemClasses, active ? activeClasses : inactiveClasses)}
                >
                  <Icon
                    name={item.icon}
                    className={cn("size-4 shrink-0", active ? "text-accent" : "text-ink-faint")}
                  />
                  {item.label}
                </Link>
              );
            })}
          </div>
        ))}
      </nav>

      <p className="px-3 text-xs leading-relaxed text-ink-faint">
        Unaudited software. Not somewhere to put money you cannot lose.
      </p>
    </div>
  );
}
