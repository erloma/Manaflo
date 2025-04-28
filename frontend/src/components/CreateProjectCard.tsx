import React, { useState, FormEvent } from "react";
import { Card } from "./ui/card";
import {
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { createProjectService } from "@/lib/api/services/projects";
import { CreateProjectRequest } from "@/lib/api/types/project";

export function CardWithForm() {
  const [name, setName] = useState<string>("");
  const [type, setType] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (!name || !type) {
      setError("Name and Type are required");
      return;
    }

    const payload : CreateProjectRequest = {
      name,
      description,
      type,
      created_by: 1, // TODO change this to real user id
      attachments: [],
    };

    try {
      const res = await createProjectService(payload); 

      const body = await res.json();
      if (!res.ok) {
        setError(body.error || "Failed to create project");
        return;
      }

      setSuccess(`Created project “${body.name}” (ID ${body.ID})`);
      setName("");
      setDescription("");
      setType("");
    } catch {
      setError("Network error");
    }
  };

  return (
    <Card className="w-[350px]">
      <CardHeader>
        <CardTitle>Create project</CardTitle>
        <CardDescription>Create a new project in 3 clicks.</CardDescription>
      </CardHeader>

      <CardContent>
        <form id="project-form" onSubmit={handleSubmit} className="space-y-4">
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="name">Name</Label>
              <Input
                id="name"
                placeholder="Name of your project"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="framework">Type</Label>
              <Select onValueChange={setType} value={type}>
                <SelectTrigger id="framework">
                  <SelectValue placeholder="Select" />
                </SelectTrigger>
                <SelectContent position="popper">
                  <SelectItem value="School project">School project</SelectItem>
                  <SelectItem value="Personal project">
                    Personal project
                  </SelectItem>
                  <SelectItem value="Work project">Work project</SelectItem>
                  <SelectItem value="Other">Other</SelectItem>
                </SelectContent>
              </Select>
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="description">Description</Label>
              <textarea
                id="description"
                placeholder="Enter project description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className="w-full rounded border p-2"
              />
            </div>

            {error && <p className="text-red-500">{error}</p>}
            {success && <p className="text-green-500">{success}</p>}
          </div>
        </form>
      </CardContent>

      <CardFooter className="flex justify-between">
        <Button
          variant="outline"
          onClick={() => {
            setName("");
            setType("");
            setDescription("");
            setError("");
            setSuccess("");
          }}
        >
          Cancel
        </Button>
        <Button type="submit" form="project-form">
          Create
        </Button>
      </CardFooter>
    </Card>
  );
}
