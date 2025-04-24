import {Button} from "./ui/button.tsx"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {jwtDecode} from "jwt-decode"
import React, { useState } from "react";

import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
  } from "@/components/ui/card"

function AccountSettings (){
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPasswordFirst, setNewPasswordFirst] = useState("");
    const [newPasswordSecond, setNewPasswordSecond] = useState("");
    const [error, setError] = useState("");

    const [editFirstName, setEditFirstName] = useState(false);
    const [editLastName, setEditLastName] = useState(false);
    const [editEmail, setEditEmail] = useState(false);
    const [editPassword, setEditPassword] = useState(false);

//----------------------------- Get User Token info ---------------------------------------------

    interface TokenPayload {
        user_id: string;
        exp: number;
        // You can add other fields here if needed
    }   

    const token = localStorage.getItem("token");

    let userID: string | null = null;

    if (token) {
        const decoded = jwtDecode<TokenPayload>(token);
        userID = decoded.user_id; // This matches the Go struct field
    }
    console.log("UserID: ", userID);

//-------------------------- Button Functions

    async function switchFirstName() {
        if (editFirstName && firstName) {
                const response = await fetch(`http://localhost:8080/api/users/${userID}`, {
                    method: "PATCH",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ firstName }),
                  });
              
                  if (!response.ok) {
                    setError(await response.text());
                    throw new Error("Error when updating first name");
                  }  
                    
                  return response.json();
        }
        setEditFirstName(!editFirstName);
    }
    async function switchLastName() {
        if (editLastName && lastName) {
                const response = await fetch(`http://localhost:8080/api/users/${userID}`, {
                    method: "PATCH",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ lastName }),
                  });
              
                  if (!response.ok) {
                    setError(await response.text());
                    throw new Error("Error when updating last name");
                  }  
                  return response.json();        
        }
        setEditLastName(!editLastName);
    }
    function switchEmail() {
        
        
        setEditEmail(!editEmail);
    }
    function switchPassword() {
        setEditPassword(!editPassword);
    }

    function updatePassword() {
        return (
            <div className="m-4">
                <div className="flex flex-col space-y-3 my-4">
                <Label htmlFor="name">Old password: </Label>
                <Input 
                    id="oldPassword"
                    value={oldPassword}
                    onChange={(e) => setOldPassword(e.target.value)} 
                    placeholder="Old Password"/>
                </div>
                <div className="flex flex-col space-y-3 my-4">
                    <Label htmlFor="name">New password </Label>
                    <Input 
                        id="newPasswordFirst" 
                        value={newPasswordFirst}
                        onChange={(e) => setNewPasswordFirst(e.target.value)} 
                        placeholder="New Password"/>
                </div>
                <div className="flex flex-col space-y-3 my-4">
                <Label htmlFor="name">Confirm new password: </Label>
                <Input 
                    id="newPasswordSecond" 
                    value={newPasswordSecond}
                    onChange={(e) => setNewPasswordSecond(e.target.value)} 
                    placeholder="New Password"/>
                <Button type="button" onClick={switchPassword} className="float-right">{"Save Password"}</Button> 
                </div>
            </div>
        )
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
                                placeholder="New first name"/> : <p>Sperma</p>}
                            {editFirstName ? <Button type="button" onClick={switchFirstName} className="float-right">Save</Button> : <img className="size-5.5"src="/images/editIcon.png" alt="edit icon" onClick={switchFirstName}/>}
                        </div>
                    </div>
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">Last Name:</Label>
                            <div className="flex items-center space-x-2 space-y-1.5">
                            {editLastName ? <Input 
                            id="lastName" 
                            value={lastName}
                            onChange={(e) => setLastName(e.target.value)} 
                            placeholder="New last name"/> : <p>Donatorn</p>}
                            {editLastName ? <Button type="button" onClick={switchLastName} className="float-right">Save</Button> : <img className="size-5.5"src="/images/editIcon.png" alt="edit icon" onClick={switchLastName}/>}
                            </div> 
                    </div> 
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">Email: </Label>
                            <div className="flex items-center space-x-2 space-y-1.5">
                                {editEmail ? <Input 
                                id="email" 
                                value={email}
                                onChange={(e) => setEmail(e.target.value)} 
                                placeholder="New e-mail"/> : <p>example@post.se</p>}
                                {editEmail ? <Button type="button" onClick={switchEmail} className="float-right">Save</Button> : <img className="size-5.5"src="/images/editIcon.png" alt="edit icon" onClick={switchEmail}/>}
                            </div>
                    </div>
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">Password: </Label>
                        <div className="flex items-center space-x-2 space-y-1.5">
                            {editPassword ? <Button type="button" onClick={switchPassword} className="float-right">{"Update Password"}</Button> : updatePassword()}
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
    )
}

export { AccountSettings }; 

