
const API_URL = "http://localhost:8080/api"


export const getUserService = async (token: string | null): Promise<Response> => {
    return fetch(`${API_URL}/user/profile`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        }
    });
};

