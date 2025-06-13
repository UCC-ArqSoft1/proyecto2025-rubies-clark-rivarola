import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './pages/Login';
import Horarios from './pages/Horarios';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/horarios" element={<Horarios />} />
      </Routes>
    </Router>
  );
}
export default App;