import { Button } from "./ui/button.tsx";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { jwtDecode } from "jwt-decode";
import { useEffect, useState } from "react";

import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import { getUserService, updateUserFieldService, updateUserPasswordService } from "@/lib/api/services/users.ts";
import { UserProfile, UpdateUserFieldRequest, UpdateUserPasswordRequest, TokenPayload } from "@/lib/api/types/user";

function AccountSettings() {

    //inputData ---------------------------
    const [firstName, setFirstName] = useState<string>("");
    const [lastName, setLastName] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [oldPassword, setOldPassword] = useState<string>("");
    const [newPasswordFirst, setNewPasswordFirst] = useState<string>("");
    const [newPasswordSecond, setNewPasswordSecond] = useState<string>("");
    const [error, setError] = useState<string>("");

    //edit states ---------------------------
    const [editFirstName, setEditFirstName] = useState<boolean>(false);
    const [editLastName, setEditLastName] = useState<boolean>(false);
    const [editEmail, setEditEmail] = useState<boolean>(false);
    const [editPassword, setEditPassword] = useState<boolean>(false);

    //stored data ---------------------------
    const [curFirstName, setCurFirstName] = useState<string>("No data");
    const [curLastName, setCurLastName] = useState<string>("No data");
    const [curEmail, setCurEmail] = useState<string>("No data");
    const [userID, setUserID] = useState<string | null>(null);

    //

    //----------------------------- Get User Token info --------------------------------------------- 

    const token = localStorage.getItem("token");

    useEffect(() => {
        if (token) {
            const decoded = jwtDecode<TokenPayload>(token);
            setUserID(decoded.user_id);
            getUser();
        }
    }, []);

    //-------------------------- Get User -----------------------------------
    async function getUser() {
        const response = await getUserService(token);

        const data: UserProfile = await response.json();
        if (data.firstName) setCurFirstName(data.firstName);
        if (data.lastName) setCurLastName(data.lastName);
        if (data.email) setCurEmail(data.email);
    }

    //-------------------------- General Update Function --------------------------------

    async function updateUserField(
        fieldName: string,
        fieldValue: string,
        currentValueSetter: React.Dispatch<React.SetStateAction<string>>,
        editSetter: React.Dispatch<React.SetStateAction<boolean>>,
        editState: boolean
    ) {
        if (editState && fieldValue && userID) {
            try {
                const updateData: UpdateUserFieldRequest = {
                    userId: userID,
                    fieldName: fieldName,
                    fieldValue: fieldValue,
                };

                const response = await updateUserFieldService(updateData);

                if (!response.ok) {
                    const data = await response.json();
                    setError(data.error);
                    throw new Error(`Error when updating ${fieldName}`);
                }

                currentValueSetter(fieldValue);
            } catch (err) {
                console.error(err);
            }
        }
        setError("");
        editSetter(!editState);
    }

    //-------------------- switch Calls --------------------
    async function switchFirstName() {
        await updateUserField("firstName", firstName, setCurFirstName, setEditFirstName, editFirstName);
    }

    async function switchLastName() {
        await updateUserField("lastName", lastName, setCurLastName, setEditLastName, editLastName);
    }

    async function switchEmail() {
        await updateUserField("email", email, setCurEmail, setEditEmail, editEmail);
    }

    //-------------------- Update Password Method --------------------
    async function switchPassword() {
        if (editPassword && oldPassword && newPasswordFirst && newPasswordSecond && userID) {
            try {
                const updatePasswordData: UpdateUserPasswordRequest = {
                    userId: userID,
                    oldPassword: oldPassword,
                    newPasswordFirst: newPasswordFirst,
                    newPasswordSecond: newPasswordSecond,
                };

                const response = await updateUserPasswordService(updatePasswordData);

                if (!response.ok) {
                    const data = await response.json();
                    setError(data.error);
                    throw new Error("Error when updating password");
                }

                alert("PASSWORD UPDATED!");
            } catch (err) {
                console.error(err);
                return;
            }
        }
        setError("");
        setEditPassword(!editPassword);
    }

    //------------------------ Update password form --------------------------------------- 
    function updatePasswordForm() {
        return (
            <div className="m-4">
                <div className="flex flex-col space-y-3 my-4">
                    <Label htmlFor="name">Old password: </Label>
                    <Input
                        id="oldPassword"
                        type="password"
                        autoComplete="current-password"
                        value={oldPassword}
                        onChange={(e) => setOldPassword(e.target.value)}
                        placeholder="Old Password" />
                </div>
                <div className="flex flex-col space-y-3 my-4">
                    <Label htmlFor="name">New password </Label>
                    <Input
                        id="newPasswordFirst"
                        type="password"
                        autoComplete="new-password"
                        value={newPasswordFirst}
                        onChange={(e) => setNewPasswordFirst(e.target.value)}
                        placeholder="New Password" />
                </div>
                <div className="flex flex-col space-y-3 my-4">
                    <Label htmlFor="name">Confirm new password: </Label>
                    <Input
                        id="newPasswordSecond"
                        type="password"
                        autoComplete="new-password"
                        value={newPasswordSecond}
                        onChange={(e) => setNewPasswordSecond(e.target.value)}
                        placeholder="New Password" />
                    <Button type="button" onClick={switchPassword} className="float-right">{"Save Password"}</Button>
                </div>
            </div>
        );
    }

    return (
        <div className="w-xl">
            <Card>
                <CardHeader>
                    <CardTitle>Account Settings</CardTitle>
                    <CardDescription>Verify or update account settings</CardDescription>
                </CardHeader>
                <CardContent className="mx-5">
                    <form>
                        <div className="grid w-full items-center gap-4">
                            <div className="flex flex-col space-y-1.5">
                                <Label htmlFor="name">First Name:</Label>
                                <div className="flex items-center space-x-2 space-y-1.5">
                                    {editFirstName ? <Input
                                        id="firstName"
                                        value={firstName}
                                        onChange={(e) => setFirstName(e.target.value)}
                                        placeholder="New first name" /> : <p>{curFirstName}</p>}
                                    {editFirstName ? <Button type="button" onClick={switchFirstName} className="float-right">Save</Button> : <img className="size-5.5" src="/images/editIcon.png" alt="edit icon" onClick={switchFirstName} />}
                                </div>
                            </div>
                            <div className="flex flex-col space-y-1.5">
                                <Label htmlFor="name">Last Name:</Label>
                                <div className="flex items-center space-x-2 space-y-1.5">
                                    {editLastName ? <Input
                                        id="lastName"
                                        value={lastName}
                                        onChange={(e) => setLastName(e.target.value)}
                                        placeholder="New last name" /> : <p>{curLastName}</p>}
                                    {editLastName ? <Button type="button" onClick={switchLastName} className="float-right">Save</Button> : <img className="size-5.5" src="/images/editIcon.png" alt="edit icon" onClick={switchLastName} />}
                                </div>
                            </div>
                            <div className="flex flex-col space-y-1.5">
                                <Label htmlFor="name">Email: </Label>
                                <div className="flex items-center space-x-2 space-y-1.5">
                                    {editEmail ? <Input
                                        id="email"
                                        value={email}
                                        onChange={(e) => setEmail(e.target.value)}
                                        placeholder="New e-mail" /> : <p>{curEmail}</p>}
                                    {editEmail ? <Button type="button" onClick={switchEmail} className="float-right">Save</Button> : <img className="size-5.5" src="/images/editIcon.png" alt="edit icon" onClick={switchEmail} />}
                                </div>
                            </div>
                            <div className="flex flex-col space-y-1.5">
                                <Label htmlFor="name">Password: </Label>
                                <div className="flex items-center space-x-2 space-y-1.5">
                                    {editPassword ? updatePasswordForm() : <Button type="button" onClick={switchPassword} className="float-right">{"Update Password"}</Button>}
                                </div>
                            </div>
                        </div>
                    </form>
                </CardContent>
                <CardFooter>
                    {error && <p className="text-red-500">{error}</p>}
                </CardFooter>
            </Card>
        </div>
    );
}

export { AccountSettings };
