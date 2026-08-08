import { feeItems } from "@/content/protocol";

export function FeeTable() {
  return (
    <div className="overflow-hidden rounded-card border border-line bg-surface">
      <div className="overflow-x-auto">
        <table className="w-full min-w-[36rem] border-collapse text-left">
          <caption className="sr-only">Every fee charged by the platform</caption>
          <thead>
            <tr className="bg-surface-muted">
              <th scope="col" className="whitespace-nowrap px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                Fee
              </th>
              <th scope="col" className="whitespace-nowrap px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                Amount
              </th>
              <th scope="col" className="px-5 py-3 text-xs font-medium uppercase tracking-[0.08em] text-ink-faint">
                When it applies
              </th>
            </tr>
          </thead>
          <tbody>
            {feeItems.map((fee) => (
              <tr key={fee.kind} className="border-t border-line align-top">
                <th scope="row" className="whitespace-nowrap px-5 py-4 text-sm font-medium text-ink">
                  {fee.label}
                </th>
                <td className="whitespace-nowrap px-5 py-4 text-sm font-semibold text-brand-ink tabular-nums">
                  {fee.value}
                </td>
                <td className="px-5 py-4 text-sm leading-relaxed text-ink-soft">{fee.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
