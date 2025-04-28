import { Task } from "../types/task";


export const createTask = async (payload: Task) : Promise<Response> => {
    return fetch("http://localhost:8081/api/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });
  };