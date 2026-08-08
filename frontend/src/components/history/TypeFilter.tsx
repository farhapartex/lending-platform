"use client";

import { ActivityKind, allTypesFilter, type TypeFilterValue } from "@/lib/enums";
import { Dropdown } from "@/components/ui/Dropdown";
import { activityLabels } from "@/components/tx/TxTypeBadge";

const options = [
  { value: allTypesFilter as TypeFilterValue, label: "All types" },
  ...Object.values(ActivityKind).map((kind) => ({
    value: kind as TypeFilterValue,
    label: activityLabels[kind],
  })),
];

type TypeFilterProps = {
  value: TypeFilterValue;
  onChange: (value: TypeFilterValue) => void;
};

export function TypeFilter({ value, onChange }: TypeFilterProps) {
  return (
    <Dropdown id="history-type-filter" label="Type" value={value} options={options} onChange={onChange} className="w-52" />
  );
}
