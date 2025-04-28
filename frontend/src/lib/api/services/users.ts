
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
    userID: string,
    fieldName: string,
    fieldValue: string
  ) : Promise<Response> => {
    return fetch(`${API_URL}/users/${userID}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ [fieldName]: fieldValue }),
    });
  };
  
  export const updateUserPasswordService = (
    userID: string,
    oldPassword: string,
    newPasswordFirst: string,
    newPasswordSecond: string
  ) : Promise<Response> => {
    return fetch(`${API_URL}/users/${userID}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ oldPassword, newPasswordFirst, newPasswordSecond }),
    });
  };
