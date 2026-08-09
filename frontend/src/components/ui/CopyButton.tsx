"use client";

import { useState } from "react";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { Button } from "@/components/ui/Button";

const feedbackDurationMs = 1600;

type CopyButtonProps = {
  value: string;
  label?: string;
  copiedLabel?: string;
};

export function CopyButton({ value, label = "Copy", copiedLabel = "Copied" }: CopyButtonProps) {
  const [copied, setCopied] = useState(false);

  const handleCopy = () => {
    void navigator.clipboard
      .writeText(value)
      .then(() => {
        setCopied(true);
        window.setTimeout(() => setCopied(false), feedbackDurationMs);
      })
      .catch(() => setCopied(false));
  };

  return (
    <Button
      variant={ButtonVariant.Ghost}
      size={ButtonSize.Sm}
      leadingIcon={copied ? IconName.Check : IconName.Receipt}
      onClick={handleCopy}
      fullWidth
    >
      {copied ? copiedLabel : label}
    </Button>
  );
}
