import { RegisterRequest, UpdateUserFieldRequest, UpdateUserPasswordRequest } from "../types/user";

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


export const updateUserFieldService = (
    data: UpdateUserFieldRequest
): Promise<Response> => {
    return fetch(`${API_URL}/users/${data.userId}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ [data.fieldName]: data.fieldValue }),
    });
};

export const updateUserPasswordService = (
    data: UpdateUserPasswordRequest
): Promise<Response> => {
    return fetch(`${API_URL}/users/${data.userId}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            oldPassword: data.oldPassword,
            newPasswordFirst: data.newPasswordFirst,
            newPasswordSecond: data.newPasswordSecond,
        }),
    });
};

  export const registerUserService = async ( data: RegisterRequest) : Promise<Response> => {
    return fetch(`${API_URL}/users`, {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify(data),
    });
  };
