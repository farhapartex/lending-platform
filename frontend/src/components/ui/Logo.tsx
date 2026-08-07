import Link from "next/link";
import { AppRoute } from "@/lib/enums";
import { cn } from "@/lib/cn";

type LogoProps = {
  className?: string;
};

export function Logo({ className }: LogoProps) {
  return (
    <Link
      href={AppRoute.Home}
      className={cn(
        "inline-flex items-center gap-2.5 rounded-pill outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas",
        className,
      )}
    >
      <span className="grid size-8 place-items-center rounded-tile bg-brand text-white">
        <svg viewBox="0 0 24 24" fill="none" className="size-5" aria-hidden="true" focusable="false">
          <path d="M6 15.5l4-4.5 3.5 3 4.5-6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
        </svg>
      </span>
      <span className="text-base font-semibold tracking-tight text-ink">Lending Platform</span>
    </Link>
  );
}
