import { cn } from "@/lib/cn";
import { welcomePageContent } from "@/content/welcome";

type StepProgressProps = {
  current: number;
  total: number;
};

export function StepProgress({ current, total }: StepProgressProps) {
  return (
    <div className="flex flex-col gap-2.5">
      <div
        role="progressbar"
        aria-label={welcomePageContent.progressLabel}
        aria-valuemin={1}
        aria-valuemax={total}
        aria-valuenow={current}
        aria-valuetext={`Step ${current} of ${total}`}
        className="flex gap-1.5"
      >
        {Array.from({ length: total }, (_, index) => (
          <span
            key={index}
            className={cn(
              "h-1.5 flex-1 rounded-pill transition-colors",
              index < current ? "bg-brand" : "bg-canvas-deep",
            )}
          />
        ))}
      </div>

      <p aria-live="polite" className="text-xs text-ink-faint tabular-nums">
        Step {current} of {total}
      </p>
    </div>
  );
}
