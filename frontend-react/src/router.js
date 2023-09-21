import React from 'react';
import { BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Login from './components/Login';
import Signup from './components/Signup';
import Dashboard from './components/Dashboard';
import CreateVM from './components/CreateVM';
import Profile from './components/Profile';
import NotFound from './components/NotFound'; // Import the NotFound component


const AppRouter = () => (
  <Router>
    <Routes>
        <Route path="/" exact element={<Dashboard/>} />
        <Route path="/login" element={<Login/>} />
        <Route path="/signup" element={<Signup/>} />
        <Route path="/dashboard" element={<Dashboard/>} />
        <Route path="/createvm" element={<CreateVM/>} />
        <Route path="/profile" element={<Profile/>} />
        <Route path="*" element={<NotFound />} />
    </Routes>
  </Router>
);

export default AppRouter;