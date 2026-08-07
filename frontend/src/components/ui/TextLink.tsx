import Link from "next/link";
import type { ReactNode } from "react";
import { AppRoute, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

const baseClasses =
  "inline-flex items-center gap-1.5 rounded-sm text-sm text-ink-soft transition-colors hover:text-brand-ink outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

type TextLinkProps = {
  href: AppRoute | string;
  children: ReactNode;
  trailingIcon?: IconName;
  className?: string;
};

export function TextLink({ href, children, trailingIcon, className }: TextLinkProps) {
  return (
    <Link href={href} className={cn(baseClasses, className)}>
      {children}
      {trailingIcon ? <Icon name={trailingIcon} className="size-3.5" /> : null}
    </Link>
  );
}
