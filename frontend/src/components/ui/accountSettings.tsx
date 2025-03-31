import {Button} from "./button.tsx"


function AccountSettings (){
    return (
        <div className='border-2 border-black rounded-2xl p-3 m-4 flow-root'> 
            <h1 className="text-black-500 text-xl font-bold mb-4 underline">Account Settings</h1>
            <h3 className="text-black-200 text-l font-bold">First Name: </h3>
            <p>cumsocket</p>

            <h3 className="text-black-200 text-l font-bold">Last Name: </h3>
            <p>Petersson</p>

            <h3 className="text-black-200 text-l font-bold">E-mail: </h3>
            <p>user@example.se</p>

            <Button className="m-3 float-right">Edit</Button>

            <div></div>
        </div> 
    )
}

export { AccountSettings }; 

