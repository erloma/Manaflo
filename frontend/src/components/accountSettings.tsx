import {Button} from "./ui/button.tsx"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
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
    const [editFirstName, setEditFirstName] = useState(false);
    const [editLastName, setEditLastName] = useState(false);
    const [editEmail, setEditEmail] = useState(false);
    const [editPassword, setEditPassword] = useState(false);

    function switchFirstName() {
        setEditFirstName(!editFirstName);
    }
    function switchLastName() {
        setEditLastName(!editLastName);
    }
    function switchEmail() {
        setEditEmail(!editEmail);
    }
    function switchPassword() {
        if(editEmail) {
            //TODO: add submit logic
        }

        setEditPassword(!editPassword);
    }

    function updatePassword() {
        return (
            <div>
                <div className="flex flex-col space-y-3 my-4">
                <Label htmlFor="name">Old password: </Label>
                <Input id="Password" placeholder="Old Password"/>
                </div>
                <div className="flex flex-col space-y-3 my-4">
                    <Label htmlFor="name">New password </Label>
                    <Input id="Password" placeholder="New Password"/>
                </div>
                <div className="flex flex-col space-y-3 my-4">
                <Label htmlFor="name">Confirm new password: </Label>
                <Input id="Password" placeholder="New Password"/>
                <Button type="button" onClick={switchPassword} className="float-right">{"Save Password"}</Button> 
                </div>
            </div>
        )
    }    

    return (
        <div>
        <Card>
            <CardHeader>
                <CardTitle>Account Settings</CardTitle>
                <CardDescription>Verify or update account settings</CardDescription>
            </CardHeader>
            <CardContent>
            <form>
                <div className="grid w-full items-center gap-4">
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">First Name:</Label>
                        <div className="flex items-center space-x-2 space-y-1.5">
                                {editFirstName ? <Input className="float-left" id="name" placeholder="First Name"/> : <p>Sperma</p>}
                                <Button type="button" onClick={switchFirstName} className="float-right">{editFirstName ? "Save Name" : "Edit"}</Button>
                            </div>
                    </div>
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">Last Name:</Label>
                            <div className="flex items-center space-x-2 space-y-1.5">
                                {editLastName ? <Input className="float-left" id="name" placeholder="Last Name"/> : <p>Donatorn</p>}
                                <Button type="button" onClick={switchLastName} className="float-right">{editLastName ? "Save Name" : "Edit"}</Button>   
                            </div> 
                    </div> 
                    <div className="flex flex-col space-y-1.5">
                        <Label htmlFor="name">Email: </Label>
                            <div className="flex items-center space-x-2 space-y-1.5">
                                {editEmail ? <Input id="email" placeholder="First Name"/> : <p>example@post.se</p>}
                                <Button type="button" onClick={switchEmail} className="float-right">{editEmail ? "Save E-mail" : "Edit"}</Button>
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
                <p>Card Footer</p>
            </CardFooter>
        </Card>
        </div> 
    )
}

export { AccountSettings }; 

