import { DocBlockKind, IconName } from "@/lib/enums";
import type { DocBlock } from "@/content/learn";
import { Alert } from "@/components/ui/Alert";
import { Icon } from "@/components/ui/Icon";
import { FaqAccordion } from "@/components/learn/FaqAccordion";
import { FeeTable } from "@/components/learn/FeeTable";
import { GlossaryList } from "@/components/learn/GlossaryList";
import { HealthTierTable } from "@/components/learn/HealthTierTable";
import { LiquidationBonusExample } from "@/components/learn/LiquidationBonusExample";

type DocBlocksProps = {
  blocks: DocBlock[];
};

export function DocBlocks({ blocks }: DocBlocksProps) {
  return (
    <div className="flex flex-col gap-5">
      {blocks.map((block, index) => (
        <DocBlockView key={`${block.kind}-${index}`} block={block} />
      ))}
    </div>
  );
}

function DocBlockView({ block }: { block: DocBlock }) {
  if (block.kind === DocBlockKind.Prose) {
    return (
      <div className="flex flex-col gap-4">
        {block.paragraphs.map((paragraph) => (
          <p key={paragraph} className="text-base leading-relaxed text-ink-soft">
            {paragraph}
          </p>
        ))}
      </div>
    );
  }

  if (block.kind === DocBlockKind.Bullets) {
    return (
      <ul className="flex flex-col gap-3">
        {block.items.map((item) => (
          <li key={item} className="flex gap-3 text-base leading-relaxed text-ink-soft">
            <Icon name={IconName.Check} className="mt-1 size-4 shrink-0 text-mint" />
            <span>{item}</span>
          </li>
        ))}
      </ul>
    );
  }

  if (block.kind === DocBlockKind.Callout) {
    return (
      <Alert title={block.title} tone={block.tone} icon={IconName.Info}>
        {block.body}
      </Alert>
    );
  }

  if (block.kind === DocBlockKind.FeeTable) {
    return <FeeTable />;
  }

  if (block.kind === DocBlockKind.BonusExample) {
    return <LiquidationBonusExample />;
  }

  if (block.kind === DocBlockKind.Faq) {
    return <FaqAccordion />;
  }

  if (block.kind === DocBlockKind.Glossary) {
    return <GlossaryList />;
  }

  return <HealthTierTable />;
}
