import { IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { welcomePageContent, welcomeSteps } from "@/content/welcome";
import { Icon } from "@/components/ui/Icon";

type StepProgressProps = {
  current: number;
  onStepSelect: (index: number) => void;
};

export function StepProgress({ current, onStepSelect }: StepProgressProps) {
  const total = welcomeSteps.length;

  return (
    <div className="flex flex-col gap-3">
      <ol className="hidden flex-wrap items-center gap-x-1 gap-y-2 sm:flex">
        {welcomeSteps.map((step, index) => {
          const isDone = index < current - 1;
          const isCurrent = index === current - 1;
          const isReachable = index <= current - 1;

          return (
            <li key={step.key} className="flex items-center gap-1">
              <button
                type="button"
                disabled={!isReachable}
                aria-current={isCurrent ? "step" : undefined}
                onClick={() => onStepSelect(index)}
                className={cn(
                  "flex items-center gap-1.5 rounded-pill px-2.5 py-1.5 text-xs font-medium transition-colors outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas",
                  isCurrent && "bg-brand-soft text-brand-ink",
                  isDone && "text-ink-soft hover:bg-surface-muted hover:text-ink",
                  !isReachable && "cursor-default text-ink-faint",
                )}
              >
                <span
                  className={cn(
                    "grid size-5 shrink-0 place-items-center rounded-pill text-[0.65rem] tabular-nums",
                    isCurrent && "bg-brand text-white",
                    isDone && "bg-mint-soft text-mint-ink",
                    !isReachable && "bg-canvas-deep text-ink-faint",
                  )}
                >
                  {isDone ? <Icon name={IconName.Check} className="size-3" /> : index + 1}
                </span>
                {step.shortLabel}
              </button>

              {index < total - 1 ? <span aria-hidden="true" className="h-px w-3 bg-line-strong" /> : null}
            </li>
          );
        })}
      </ol>

      <div className="flex flex-col gap-2 sm:hidden">
        <div
          role="progressbar"
          aria-label={welcomePageContent.progressLabel}
          aria-valuemin={1}
          aria-valuemax={total}
          aria-valuenow={current}
          aria-valuetext={`Step ${current} of ${total}`}
          className="flex gap-1.5"
        >
          {welcomeSteps.map((step, index) => (
            <span
              key={step.key}
              className={cn("h-1.5 flex-1 rounded-pill transition-colors", index < current ? "bg-brand" : "bg-canvas-deep")}
            />
          ))}
        </div>
      </div>

      <p aria-live="polite" className="text-xs text-ink-faint tabular-nums">
        Step {current} of {total}
      </p>
    </div>
  );
}
