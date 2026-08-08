"use client";

import { useState } from "react";
import { BadgeTone, ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { formatDateTimeUtc } from "@/lib/format";
import type { LiquidationEvent } from "@/content/dashboard";
import { Alert } from "@/components/ui/Alert";
import { Button } from "@/components/ui/Button";
import { Drawer } from "@/components/ui/Drawer";
import { LiquidationReceipt } from "@/components/dashboard/LiquidationReceipt";

type LiquidationEventBannerProps = {
  event: LiquidationEvent;
};

export function LiquidationEventBanner({ event }: LiquidationEventBannerProps) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <Alert title="Part of your position was liquidated" tone={BadgeTone.Critical} icon={IconName.Warning}>
        <div className="flex flex-col gap-3">
          <span>
            On {formatDateTimeUtc(event.timestamp)} your health factor fell below 1.00 and a liquidator repaid part of
            your loan in exchange for some of your collateral. Here is exactly what happened and why.
          </span>
          <div>
            <Button variant={ButtonVariant.Secondary} size={ButtonSize.Sm} onClick={() => setIsOpen(true)}>
              View the details
            </Button>
          </div>
        </div>
      </Alert>

      <Drawer open={isOpen} onClose={() => setIsOpen(false)} title="Liquidation details">
        <LiquidationReceipt event={event} />
      </Drawer>
    </>
  );
}
