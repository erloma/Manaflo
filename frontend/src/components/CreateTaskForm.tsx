import { useState } from 'react';
import { DatePicker } from "./DatePicker"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"

export function CreateTaskForm() {

  const [title, setTitle] = useState("");
  const [description, setDesc] = useState("");
  const [priority, setPriority] = useState("");
  const [date, setDate] = useState<Date>(new Date);

  const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const processedTitle = event.target.value.trim(); // removes leading/trailing spaces
    setTitle(processedTitle);
  }

  const handleDescChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const processedDesc = event.target.value.trim();
    setDesc(processedDesc);
  }

  const handlePrioChange = (value: string) => {
    setPriority(value);
  }

  const handleDateChange = (selectedDate: Date ) => {
    setDate(selectedDate);
    console.log(selectedDate);
  }




  return (
    <Card className="w-[800px]">
      <CardHeader className="flex flex-col items-center text-center">
        <CardTitle>Create task</CardTitle>


        <CardDescription>Create a task for your project</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div className="grid w-full items-center gap-4">
            <div className="flex justify-between gap-x-8">
              <div className="flex flex-col w-2/3 space-y-1.5">
                <Label htmlFor="title">Title</Label>
                <Input id="title" value={title} placeholder="Title of the task" onChange={handleTitleChange} />
                <Label htmlFor="desc">Description</Label>
                <Input className="h-full" id="desc" value={description} onChange={handleDescChange} placeholder="Description of the task" />
              </div>
              <div className="flex flex-col space-y-1.5 ml-auto">
                <Label htmlFor="name">Due date</Label>
                <div className="w-200px">
                  <DatePicker date={date} onDateChange={handleDateChange} />
                </div>

              </div>
            </div>

            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="priority">Priority</Label>
              <Select value={priority} onValueChange={handlePrioChange}>
                <SelectTrigger id="priority">
                  <SelectValue placeholder="Select" />
                </SelectTrigger>
                <SelectContent position="popper">
                  <SelectItem value="low">Low</SelectItem>
                  <SelectItem value="medium">Medium</SelectItem>
                  <SelectItem value="high">High</SelectItem>
                </SelectContent>
              </Select>
              <Label htmlFor="assignee">Assignee</Label>
              <Select>
                <SelectTrigger id="assignee">
                  <SelectValue placeholder="Unassigned" />
                </SelectTrigger>
                <SelectContent position="popper">
                  {/* TODO: dynamically load project members */}
                </SelectContent>
              </Select>
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter className="flex justify-between">
        <Button variant="outline">Cancel</Button>
        <Button>Create task</Button>
      </CardFooter>
    </Card>
  )
}


export default CreateTaskForm
