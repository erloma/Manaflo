import { Calendar } from "@/components/ui/calendar"
import React from "react"

 
export function DatePicker() {
  const [date, setDate] = React.useState<Date | undefined>(new Date())

  let currentDate = new Date();
  currentDate.setHours(0,0,0,0);

 
  return (
    <Calendar
      mode="single"
      selected={date}
      onSelect={setDate}
      disabled={{before: currentDate}}
      className="rounded-md border shadow"
    />
  )
}