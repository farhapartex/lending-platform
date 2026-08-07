import { AmountValidationCode, IconName } from "@/lib/enums";
import { Icon } from "@/components/ui/Icon";

type AmountValidationMessageProps = {
  id: string;
  code: AmountValidationCode;
  messages: Record<AmountValidationCode, string | null>;
};

export function AmountValidationMessage({ id, code, messages }: AmountValidationMessageProps) {
  const message = messages[code];

  if (message === null) {
    return <span id={id} className="sr-only" />;
  }

  return (
    <p id={id} role="alert" className="flex items-start gap-1.5 text-sm text-rose-ink">
      <Icon name={IconName.Warning} className="mt-0.5 size-4" />
      {message}
    </p>
  );
}
