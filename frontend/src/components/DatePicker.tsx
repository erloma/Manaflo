import { Calendar } from "@/components/ui/calendar";
import React from "react";

interface DatePickerProps {
  date: Date;
  onDateChange: (date: Date) => void;
}

export function DatePicker({ date, onDateChange }: DatePickerProps) {

  let currentDate = new Date();
  currentDate.setHours(0, 0, 0, 0);

  const handleSelect = (selectedDate: Date | undefined) => {
    if (selectedDate) {
      onDateChange(selectedDate);
    }
  };


  return (
    <Calendar
      mode="single"
      selected={date || undefined} 
      onSelect={handleSelect}
      disabled={{ before: currentDate }}
      className="rounded-md border shadow"
    />
  );
}

export default DatePicker;
