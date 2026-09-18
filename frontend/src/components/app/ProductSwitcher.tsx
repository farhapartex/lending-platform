"use client";

import Link from "next/link";
import { cn } from "@/lib/cn";
import { navProducts, type NavProduct } from "@/content/navigation";

const baseClasses =
  "rounded-pill px-3 py-1.5 text-sm transition-colors outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas";

type ProductSwitcherProps = {
  active: NavProduct;
  onNavigate?: () => void;
  className?: string;
};

export function ProductSwitcher({ active, onNavigate, className }: ProductSwitcherProps) {
  return (
    <div
      role="group"
      aria-label="Choose a product"
      className={cn("flex items-center gap-1 rounded-pill border border-line bg-surface-muted p-1", className)}
    >
      {navProducts.map((product) => {
        const isActive = product.key === active.key;

        return (
          <Link
            key={product.key}
            href={product.home}
            aria-current={isActive ? "true" : undefined}
            onClick={onNavigate}
            className={cn(
              baseClasses,
              isActive ? "bg-surface font-medium text-ink shadow-soft" : "text-ink-soft hover:text-ink",
            )}
          >
            {product.label}
          </Link>
        );
      })}
    </div>
  );
}
