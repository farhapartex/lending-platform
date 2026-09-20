import { BadgeTone, ButtonSize, ButtonVariant, IconName, SurfaceElevation } from "@/lib/enums";
import type { ProductSummary } from "@/content/landing";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Card } from "@/components/ui/Card";
import { Icon } from "@/components/ui/Icon";

type ProductCardProps = {
  product: ProductSummary;
  accessory?: React.ReactNode;
};

export function ProductCard({ product, accessory }: ProductCardProps) {
  return (
    <Card
      elevation={product.isLive ? SurfaceElevation.Lifted : SurfaceElevation.Flat}
      className={`flex h-full flex-col gap-5 p-6 sm:p-7 ${product.isLive ? "" : "border-dashed"}`}
    >
      <div className="flex flex-wrap items-center justify-between gap-3">
        <h3 className="text-xl font-semibold tracking-tight text-ink">{product.title}</h3>
        <Badge tone={product.isLive ? BadgeTone.Positive : BadgeTone.Caution}>{product.eyebrow}</Badge>
      </div>

      <p className="text-sm leading-relaxed text-ink-soft">{product.description}</p>

      {accessory}

      <ul className="flex flex-col gap-2.5">
        {product.points.map((point) => (
          <li key={point} className="flex items-start gap-2.5 text-sm leading-relaxed text-ink-soft">
            <Icon name={IconName.Check} className="mt-0.5 size-4 shrink-0 text-mint" />
            {point}
          </li>
        ))}
      </ul>

      <div className="mt-auto pt-2">
        <Button
          href={product.href}
          size={ButtonSize.Md}
          variant={product.isLive ? ButtonVariant.Primary : ButtonVariant.Secondary}
          trailingIcon={IconName.ArrowRight}
        >
          {product.cta}
        </Button>
      </div>
    </Card>
  );
}
