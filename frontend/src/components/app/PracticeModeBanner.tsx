import { AppRoute, IconName, NetworkKind } from "@/lib/enums";
import { activeNetwork } from "@/content/protocol";
import { Container } from "@/components/ui/Container";
import { Icon } from "@/components/ui/Icon";
import { TextLink } from "@/components/ui/TextLink";

export function PracticeModeBanner() {
  if (activeNetwork !== NetworkKind.Testnet) {
    return null;
  }

  return (
    <div className="border-b border-brand-border bg-brand-soft">
      <Container className="flex flex-wrap items-center justify-center gap-x-3 gap-y-1 py-2 text-center">
        <span className="flex items-center gap-2 text-xs font-medium text-brand-ink">
          <Icon name={IconName.Beaker} className="size-3.5" />
          Practice mode: you are on a test network and these tokens have no value.
        </span>
        <TextLink href={AppRoute.Practice} className="text-xs text-brand-ink">
          What this means
        </TextLink>
      </Container>
    </div>
  );
}
