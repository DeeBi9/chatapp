import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Auth from './auth/Auth';
import Welcome from './LandingPage/Welcome';
import Signin from './signin/Signin';
import MainChat from './main/Main';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Welcome/>}></Route>
        <Route path="/auth" element={<Auth/>}></Route>
        <Route path='/signin/' element={<Signin/>}></Route>  
        <Route path='/main/:userId' element={<MainChat/>}></Route>
      </Routes> 

    </BrowserRouter>
  );
}

export default App;
