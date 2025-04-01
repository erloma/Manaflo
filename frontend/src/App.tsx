
import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "./pages/HomePage";
import FirstPage from "./pages/FirstPage"
import LoginPage from "./pages/LoginPage.tsx";

function App() {
  return (
    <BrowserRouter>
    <Routes>
      <Route path="/" element={<HomePage />}>
      </Route>
      <Route path="first" element={<FirstPage />} />
      <Route path="login" element={<LoginPage />} />
    </Routes>
  </BrowserRouter>
  );
}

export default App;
