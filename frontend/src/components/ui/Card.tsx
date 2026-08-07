import type { ReactNode } from "react";
import { SurfaceElevation } from "@/lib/enums";
import { cn } from "@/lib/cn";

const elevationClasses: Record<SurfaceElevation, string> = {
  [SurfaceElevation.Flat]: "border-line",
  [SurfaceElevation.Raised]: "border-line shadow-soft",
  [SurfaceElevation.Lifted]: "border-line-strong shadow-lift",
};

type CardProps = {
  children: ReactNode;
  elevation?: SurfaceElevation;
  interactive?: boolean;
  className?: string;
};

export function Card({ children, elevation = SurfaceElevation.Raised, interactive = false, className }: CardProps) {
  return (
    <div
      className={cn(
        "rounded-card border bg-surface",
        elevationClasses[elevation],
        interactive && "transition-shadow duration-200 hover:shadow-lift",
        className,
      )}
    >
      {children}
    </div>
  );
}
