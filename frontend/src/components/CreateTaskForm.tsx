import * as React from "react"

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
  return (
    <Card className="w-[700px]">
      <CardHeader className="flex flex-col items-center text-center">
          <CardTitle>Create task</CardTitle>

        
        <CardDescription>Create a task for your project</CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <div className="grid w-full items-center gap-4">
            <div className="flex justify-between">
              <div className="flex flex-col w-2/3 space-y-1.5">
                <Label htmlFor="title">Title</Label>
                <Input id="title" placeholder="Title of the task" />
                <Label htmlFor="desc">Description</Label>
                <Input className="h-full" id="desc" placeholder="Description of the task" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Title</Label>
                <Input id="name" placeholder="Title of the task" />
                <Label htmlFor="name">Title</Label>
                <Input id="name" placeholder="Title of the task" />
                <Label htmlFor="name">Title</Label>
                <Input id="name" placeholder="Title of the task" />
              </div>
            </div>

            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="priority">Priority</Label>
              <Select>
                <SelectTrigger id="priority">
                  <SelectValue placeholder="Select" />
                </SelectTrigger>
                <SelectContent position="popper">
                  <SelectItem value="low">Low</SelectItem>
                  <SelectItem value="medium">Medium</SelectItem>
                  <SelectItem value="high">High</SelectItem>
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