import { LoginRequest } from "../types/user";

const API_URL = "http://localhost:8080/api"


export const loginUserService = async (data: LoginRequest) : Promise<Response> => {
    return fetch(`${API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    });
};