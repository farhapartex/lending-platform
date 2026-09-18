import { BadgeTone, IconName } from "@/lib/enums";
import { fusdPreviewNotice } from "@/content/fusdPreview";
import { Alert } from "@/components/ui/Alert";
import { Container } from "@/components/ui/Container";

export function FusdPreviewNotice() {
  return (
    <Container className="pt-6">
      <Alert title={fusdPreviewNotice.title} tone={BadgeTone.Caution} icon={IconName.Warning}>
        {fusdPreviewNotice.body}
      </Alert>
    </Container>
  );
}
