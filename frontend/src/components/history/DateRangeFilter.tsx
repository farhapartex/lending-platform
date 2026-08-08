"use client";

import { DateRangePreset } from "@/lib/enums";
import { Dropdown } from "@/components/ui/Dropdown";

const options = [
  { value: DateRangePreset.AllTime, label: "All time" },
  { value: DateRangePreset.Last7Days, label: "Last 7 days" },
  { value: DateRangePreset.Last30Days, label: "Last 30 days" },
  { value: DateRangePreset.Last90Days, label: "Last 90 days" },
];

type DateRangeFilterProps = {
  value: DateRangePreset;
  onChange: (value: DateRangePreset) => void;
};

export function DateRangeFilter({ value, onChange }: DateRangeFilterProps) {
  return (
    <Dropdown
      id="history-date-filter"
      label="Period"
      value={value}
      options={options}
      onChange={onChange}
      className="w-44"
    />
  );
}
