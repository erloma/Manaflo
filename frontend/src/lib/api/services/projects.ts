import { CreateProjectRequest } from "../types/project";
import { UserInfo } from "../types/user";
const API_URL = "http://localhost:8082/api"

export const createProjectService = (data: CreateProjectRequest): Promise<Response> => {
    return fetch(`${API_URL}/projects`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
};

export const getUsersInProjectService = async (project_id: string, token: string | null): Promise<UserInfo[]> => {
    const response = await fetch(`${API_URL}/projects/${project_id}/users`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
    });

    if (!response.ok) {
        throw new Error("Failed to fetch users in project");
    }

    const users: UserInfo[] = await response.json();
    return users;
};

