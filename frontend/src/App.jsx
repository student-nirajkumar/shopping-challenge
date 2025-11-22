import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import ItemsList from "./pages/ItemsList";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/items" element={<ItemsList />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
