import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import FirstPage from "./pages/FirstPage"
import LoginPage from "./pages/LoginPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";

function App() {
  return (
    <BrowserRouter>
    <Routes>
      <Route path="/home" element={<HomePage />}>
      </Route>
      <Route path="first" element={<FirstPage />} />
      <Route path="login" element={<LoginPage />} />
      <Route path="register" element={<RegisterPage />} />
    </Routes>
  </BrowserRouter>
  );
}

export default App;
