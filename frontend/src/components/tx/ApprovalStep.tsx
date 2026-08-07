import { IconName, StepState } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

export type ApprovalStepItem = {
  label: string;
  description: string;
  state: StepState;
};

const markerClasses: Record<StepState, string> = {
  [StepState.Done]: "border-mint-border bg-mint-soft text-mint-ink",
  [StepState.Active]: "border-brand bg-brand text-white",
  [StepState.Upcoming]: "border-line-strong bg-surface text-ink-faint",
};

const labelClasses: Record<StepState, string> = {
  [StepState.Done]: "text-ink",
  [StepState.Active]: "text-ink",
  [StepState.Upcoming]: "text-ink-soft",
};

type ApprovalStepProps = {
  steps: ApprovalStepItem[];
};

export function ApprovalStep({ steps }: ApprovalStepProps) {
  return (
    <ol className="flex flex-col gap-3">
      {steps.map((step, index) => (
        <li key={step.label} className="flex gap-3">
          <span
            className={cn(
              "grid size-7 shrink-0 place-items-center rounded-pill border text-xs font-semibold tabular-nums",
              markerClasses[step.state],
            )}
          >
            {step.state === StepState.Done ? <Icon name={IconName.Check} className="size-3.5" /> : index + 1}
          </span>
          <div className="flex flex-col gap-0.5">
            <span className={cn("text-sm font-medium", labelClasses[step.state])}>{step.label}</span>
            <span className="text-xs leading-relaxed text-ink-soft">{step.description}</span>
          </div>
        </li>
      ))}
    </ol>
  );
}
