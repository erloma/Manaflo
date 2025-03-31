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

    function switchFirstName() {
        setEditFirstName(!editFirstName);
    }
    function switchLastName() {
        setEditLastName(!editLastName);
    }
    function switchEmail() {
        setEditEmail(!editEmail);
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
                </div>
            </form>
            </CardContent>

            <CardFooter>
                <p>Card Footer</p>
            </CardFooter>
        </Card>
        </div>

        /* <div className='border-2 border-black rounded-2xl p-3 m-4 flow-root max-w-xl'> 
            <h1 className="text-black-500 text-xl font-bold mb-4 underline">Account Settings</h1>
            <h3 className="text-black-200 text-l font-bold">First Name: </h3>
            <p>cumsocket</p>

            <h3 className="text-black-200 text-l font-bold">Last Name: </h3>
            <div className="flow-root p-0.5">
                <p>Petersson</p>
                <Button className="float-right">Edit</Button>
            </div>

            <h3 className="text-black-200 text-l font-bold">E-mail: </h3>
            <div className="flow-root m-auto">
                <p className="my-auto float-left">user@example.se</p> 
                <Button className="ml-5">Edit</Button>
            </div>
            
            
            <h3 className="text-black-200 text-l font-bold">Password: </h3>
            <p>user@example.se</p>


            

            <div></div>
        </div> */  
    )
}

export { AccountSettings }; 

