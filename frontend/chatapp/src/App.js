import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Auth from './auth/Auth';
import Welcome from './LandingPage/Welcome';
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Welcome/>}></Route>
        <Route path="/auth" element={<Auth/>}>
        </Route>
      </Routes>

    </BrowserRouter>
  );
}

export default App;
