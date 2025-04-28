import { CreateProjectRequest } from "../types/project";
const API_URL = "http://localhost:8082/api"

export const createProjectService = (data: CreateProjectRequest) : Promise<Response> => {
    return fetch(`${API_URL}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
};