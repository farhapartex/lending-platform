import type { ReactNode } from "react";
import { IconName } from "@/lib/enums";
import { cn } from "@/lib/cn";

const glyphs: Record<IconName, ReactNode> = {
  [IconName.ArrowRight]: (
    <>
      <path d="M5 12h14" />
      <path d="M13 6l6 6-6 6" />
    </>
  ),
  [IconName.ArrowUpRight]: (
    <>
      <path d="M7 17L17 7" />
      <path d="M8.5 7H17v8.5" />
    </>
  ),
  [IconName.Wallet]: (
    <>
      <rect x="3" y="6" width="18" height="13" rx="2.5" />
      <path d="M3 10.5h18" />
      <path d="M16.5 14.75h.01" />
    </>
  ),
  [IconName.Coins]: (
    <>
      <circle cx="9.5" cy="9.5" r="5.5" />
      <path d="M14.4 6.7A5.5 5.5 0 1 1 11 19.9" />
    </>
  ),
  [IconName.ShieldCheck]: (
    <>
      <path d="M12 3l7.5 3v5.9c0 4.2-3 7.4-7.5 9.1-4.5-1.7-7.5-4.9-7.5-9.1V6z" />
      <path d="M9 12l2.2 2.2L15.5 10" />
    </>
  ),
  [IconName.Gauge]: (
    <>
      <path d="M4.2 18a9 9 0 1 1 15.6 0" />
      <path d="M12 14.5l3.5-3.8" />
    </>
  ),
  [IconName.TrendUp]: (
    <>
      <path d="M3 17l6-6 4 4 7-8" />
      <path d="M16 7h4v4" />
    </>
  ),
  [IconName.TrendDown]: (
    <>
      <path d="M3 7l6 6 4-4 7 8" />
      <path d="M16 17h4v-4" />
    </>
  ),
  [IconName.Minus]: <path d="M6 12h12" />,
  [IconName.Check]: <path d="M5 12.5l4.5 4.5L19 7.5" />,
  [IconName.Code]: (
    <>
      <path d="M9 8l-4 4 4 4" />
      <path d="M15 8l4 4-4 4" />
    </>
  ),
  [IconName.Lock]: (
    <>
      <rect x="4.5" y="10.5" width="15" height="9.5" rx="2.5" />
      <path d="M8.5 10.5V8a3.5 3.5 0 1 1 7 0v2.5" />
    </>
  ),
  [IconName.Sliders]: (
    <>
      <path d="M4 8.5h9" />
      <path d="M18.5 8.5H20" />
      <path d="M4 15.5h4" />
      <path d="M13.5 15.5H20" />
      <circle cx="15.75" cy="8.5" r="2.25" />
      <circle cx="10.25" cy="15.5" r="2.25" />
    </>
  ),
  [IconName.Beaker]: (
    <>
      <path d="M8.5 3h7" />
      <path d="M10 3v5.4l-4.2 8.4A2.2 2.2 0 0 0 7.8 20h8.4a2.2 2.2 0 0 0 2-3.2L14 8.4V3" />
      <path d="M7.6 14.5h8.8" />
    </>
  ),
  [IconName.Menu]: (
    <>
      <path d="M4 7h16" />
      <path d="M4 12h16" />
      <path d="M4 17h16" />
    </>
  ),
  [IconName.Close]: (
    <>
      <path d="M6.5 6.5l11 11" />
      <path d="M17.5 6.5l-11 11" />
    </>
  ),
  [IconName.ExternalLink]: (
    <>
      <path d="M14 4h6v6" />
      <path d="M20 4l-8.5 8.5" />
      <path d="M18 14.5V18a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3.5" />
    </>
  ),
  [IconName.Receipt]: (
    <>
      <path d="M6 3h12v18l-3-1.8-3 1.8-3-1.8L6 21z" />
      <path d="M9.5 8.5h5" />
      <path d="M9.5 12.5h5" />
    </>
  ),
  [IconName.Info]: (
    <>
      <circle cx="12" cy="12" r="8.5" />
      <path d="M12 11v5.5" />
      <path d="M12 7.75h.01" />
    </>
  ),
  [IconName.Warning]: (
    <>
      <path d="M12 4.5l8.5 15H3.5z" />
      <path d="M12 10v4" />
      <path d="M12 16.75h.01" />
    </>
  ),
  [IconName.Loader]: (
    <>
      <path d="M12 3.5v3.5" />
      <path d="M12 17v3.5" />
      <path d="M3.5 12H7" />
      <path d="M17 12h3.5" />
      <path d="M6.2 6.2l2.4 2.4" />
      <path d="M15.4 15.4l2.4 2.4" />
      <path d="M6.2 17.8l2.4-2.4" />
      <path d="M15.4 8.6l2.4-2.4" />
    </>
  ),
};

type IconProps = {
  name: IconName;
  className?: string;
};

export function Icon({ name, className }: IconProps) {
  return (
    <svg
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth={1.75}
      strokeLinecap="round"
      strokeLinejoin="round"
      aria-hidden="true"
      focusable="false"
      className={cn("shrink-0", className)}
    >
      {glyphs[name]}
    </svg>
  );
}
