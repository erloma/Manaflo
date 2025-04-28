import { useState, FormEvent } from "react";
import { Card } from "./ui/card";
import {
  CardContent,
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

export function CardWithForm() {
  const [name, setName] = useState<string>("");
  const [type, setType] = useState<string>("");
  const [desc, setDescription] = useState<string>("");
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

    const payload = {
      name,
      description: desc,
      created_by: 1, // TODO change this to real user id
      attachments: [],
    };

    try {
      const res = await fetch("/api/projects", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      const body = await res.json();
      if (!res.ok) {
        setError(body.error || "Failed to create project");
        return;
      }

      setSuccess(`Created project “${body.name}” (ID ${body.ID})`);
      setName("");
      setType("");
      setDescription(""); 
    } catch {
      setError("Network error");
    }
  };

  return (
    <Card className="w-[350px]">
      <CardHeader>
        <CardTitle>Create project</CardTitle>
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
              <Label htmlFor="description">Description</Label>
              <Input
                id="description"
                placeholder="Describe your project"
                value={desc}
                onChange={(e) => setDescription(e.target.value)}
                className="h-24 resize-none" 
              />
            </div>

            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="type">Type</Label>
              <Select onValueChange={setType} value={type}>
                <SelectTrigger id="type">
                  <SelectValue placeholder="Select" />
                </SelectTrigger>
                <SelectContent position="popper">
                  <SelectItem value="school">School project</SelectItem>
                  <SelectItem value="personal">Personal project</SelectItem>
                  <SelectItem value="work">Work project</SelectItem>
                </SelectContent>
              </Select>
            </div>

            

            {error && <p className="text-red-500">{error}</p>}
            {success && <p className="text-green-500">{success}</p>}
          </div>
        </form>
      </CardContent>

      <CardFooter className="flex justify-between">
        <Button variant="outline">Cancel</Button>
        <Button type="submit" form="project-form">
          Create
        </Button>
      </CardFooter>
    </Card>
  );
}
