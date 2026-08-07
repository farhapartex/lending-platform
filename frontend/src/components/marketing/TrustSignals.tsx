import { BadgeTone, SectionId, SectionTone, TrustSignalStatus } from "@/lib/enums";
import { trustSignals } from "@/content/landing";
import { Badge } from "@/components/ui/Badge";
import { Icon } from "@/components/ui/Icon";
import { Section } from "@/components/ui/Section";
import { SectionHeading } from "@/components/ui/SectionHeading";

const statusLabels: Record<TrustSignalStatus, string> = {
  [TrustSignalStatus.Live]: "In place",
  [TrustSignalStatus.Planned]: "Planned",
};

const statusTones: Record<TrustSignalStatus, BadgeTone> = {
  [TrustSignalStatus.Live]: BadgeTone.Positive,
  [TrustSignalStatus.Planned]: BadgeTone.Caution,
};

export function TrustSignals() {
  return (
    <Section id={SectionId.Trust} tone={SectionTone.Canvas}>
      <SectionHeading
        sectionId={SectionId.Trust}
        eyebrow="Security"
        title="What protects you, stated plainly."
        description="We label what is already in place and what is still ahead, because a platform holding funds should not blur the difference."
      />

      <ul className="mt-10 grid gap-5 sm:grid-cols-2">
        {trustSignals.map((signal) => (
          <li
            key={signal.key}
            className="flex gap-4 rounded-card border border-line bg-surface p-6 shadow-soft"
          >
            <span className="grid size-10 shrink-0 place-items-center rounded-tile bg-surface-muted text-ink-soft">
              <Icon name={signal.icon} className="size-5" />
            </span>
            <div className="flex flex-col gap-2">
              <div className="flex flex-wrap items-center gap-2.5">
                <h3 className="text-base font-semibold tracking-tight text-ink">{signal.title}</h3>
                <Badge tone={statusTones[signal.status]}>{statusLabels[signal.status]}</Badge>
              </div>
              <p className="text-sm leading-relaxed text-ink-soft">{signal.description}</p>
            </div>
          </li>
        ))}
      </ul>
    </Section>
  );
}
