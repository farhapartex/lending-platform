import Link from "next/link";
import type { ReactNode } from "react";
import { AppRoute, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";
import { Icon } from "@/components/ui/Icon";

const baseClasses =
  "inline-flex items-center justify-center rounded-pill font-medium transition-colors duration-150 outline-none focus-visible:ring-2 focus-visible:ring-brand focus-visible:ring-offset-2 focus-visible:ring-offset-canvas disabled:cursor-not-allowed disabled:opacity-55";

const variantClasses: Record<ButtonVariant, string> = {
  [ButtonVariant.Primary]: "bg-brand text-white shadow-soft hover:bg-brand-hover active:bg-brand-active",
  [ButtonVariant.Secondary]: "border border-line-strong bg-surface text-ink hover:border-brand-border hover:text-brand-ink",
  [ButtonVariant.Subtle]: "bg-brand-soft text-brand-ink hover:bg-brand-muted",
  [ButtonVariant.Ghost]: "text-ink-soft hover:bg-surface-muted hover:text-ink",
};

const sizeClasses: Record<ButtonSize, string> = {
  [ButtonSize.Sm]: "h-9 gap-1.5 px-3.5 text-sm",
  [ButtonSize.Md]: "h-11 gap-2 px-5 text-sm",
  [ButtonSize.Lg]: "h-12 gap-2 px-6 text-base",
};

const iconSizeClasses: Record<ButtonSize, string> = {
  [ButtonSize.Sm]: "size-4",
  [ButtonSize.Md]: "size-4",
  [ButtonSize.Lg]: "size-[1.125rem]",
};

type ButtonCommonProps = {
  children: ReactNode;
  variant?: ButtonVariant;
  size?: ButtonSize;
  leadingIcon?: IconName;
  trailingIcon?: IconName;
  fullWidth?: boolean;
  className?: string;
};

type ButtonRouteProps = ButtonCommonProps & {
  href: AppRoute | string;
};

type ButtonActionProps = ButtonCommonProps & {
  href?: never;
  onClick?: () => void;
  disabled?: boolean;
  ariaLabel?: string;
  ariaExpanded?: boolean;
  ariaControls?: string;
};

export type ButtonProps = ButtonRouteProps | ButtonActionProps;

export function Button(props: ButtonProps) {
  const {
    children,
    variant = ButtonVariant.Primary,
    size = ButtonSize.Md,
    leadingIcon,
    trailingIcon,
    fullWidth = false,
    className,
  } = props;

  const classes = cn(baseClasses, variantClasses[variant], sizeClasses[size], fullWidth && "w-full", className);

  const content = (
    <>
      {leadingIcon ? <Icon name={leadingIcon} className={iconSizeClasses[size]} /> : null}
      <span>{children}</span>
      {trailingIcon ? <Icon name={trailingIcon} className={iconSizeClasses[size]} /> : null}
    </>
  );

  if (props.href !== undefined) {
    return (
      <Link href={props.href} className={classes}>
        {content}
      </Link>
    );
  }

  return (
    <button
      type="button"
      className={classes}
      onClick={props.onClick}
      disabled={props.disabled}
      aria-label={props.ariaLabel}
      aria-expanded={props.ariaExpanded}
      aria-controls={props.ariaControls}
    >
      {content}
    </button>
  );
}
