import { Task } from "../types/task";


export const createTask = async (payload: Task, token: string | null): Promise<Response> => {
  return fetch("http://localhost:8081/api/tasks", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${token}`
    },
    body: JSON.stringify(payload),
  });
};
