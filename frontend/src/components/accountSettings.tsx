import {Button} from "./ui/button.tsx"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
  } from "@/components/ui/card"

function AccountSettings (){
    var editName = true;

    function nameEdit() {
        if (editName) {
            return <Input id="name" placeholder="First Name" />
        } else {
            return <p>Sperma Doneraren</p>
        }
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
                        {nameEdit()}
                    </div>
                    <div className="flex flex-col space-y-1.5">
                    <Label htmlFor="name">Last Name:</Label>
                        {nameEdit()}
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

