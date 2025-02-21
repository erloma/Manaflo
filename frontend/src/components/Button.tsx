import { useState } from "react";

const Button = () => {
  const [responseMessage, setResponseMessage] = useState("");

  const handleClick = async () => {
    try {
      const res = await fetch("http://localhost:8080/ping");
      const data = await res.text();
      setResponseMessage(data);
    } catch (error) {
      console.error("Error fetching data");
      setResponseMessage("Error conecting backend");
    }
  };

  return (
    <div className="flex flex-col items-center">
      <button className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-700 transition" onClick={handleClick}>
        Click Me
      </button>
      {responseMessage && <p className="mt-4">{responseMessage}</p>}
    </div>
  );
};

export default Button;