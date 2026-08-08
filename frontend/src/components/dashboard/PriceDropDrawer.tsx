"use client";

import { useState } from "react";
import { ButtonSize, ButtonVariant, IconName } from "@/lib/enums";
import { borrowPageContent } from "@/content/borrow";
import { dashboardContent } from "@/content/dashboard";
import { Button } from "@/components/ui/Button";
import { Drawer } from "@/components/ui/Drawer";
import { PriceDropSimulatorBody } from "@/components/borrow/PriceDropSimulatorBody";

export function PriceDropDrawer() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <Button
        variant={ButtonVariant.Subtle}
        size={ButtonSize.Sm}
        leadingIcon={IconName.Sliders}
        onClick={() => setIsOpen(true)}
      >
        {dashboardContent.simulatorTrigger}
      </Button>

      <Drawer open={isOpen} onClose={() => setIsOpen(false)} title={borrowPageContent.simulatorTitle}>
        <div className="flex flex-col gap-4">
          <p className="text-sm leading-relaxed text-ink-soft">{borrowPageContent.simulatorDescription}</p>
          <PriceDropSimulatorBody />
        </div>
      </Drawer>
    </>
  );
}
