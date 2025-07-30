import { BrowserRouter, Routes, Route } from "react-router-dom";
import Card from "./components/Card"
import HomePage from "./components/HomePage";
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/card" element={<Card />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
