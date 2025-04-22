import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import FirstPage from "./pages/FirstPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />}></Route>
        <Route path="first" element={<FirstPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
