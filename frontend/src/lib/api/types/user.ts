export interface RegisterRequest {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
}

export interface TokenPayload {
    user_id: string;
    exp: number;
}

export interface UserProfile {
    firstName: string;
    lastName: string;
    email: string;
}

export interface UpdateUserFieldRequest {
    userId: string;
    fieldName: string;
    fieldValue: string;
}

export interface UpdateUserPasswordRequest {
    userId: string;
    oldPassword: string;
    newPasswordFirst: string;
    newPasswordSecond: string;
}