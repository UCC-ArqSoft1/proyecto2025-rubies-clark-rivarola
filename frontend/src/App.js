import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import Home from './views/Home';
import Login from './views/Login';
import DetalleActividad from './views/DetalleActividad';
import MisActividades from './views/MisActividades';
import NuevaActividad from './views/NuevaActividad';

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/actividad/:id" element={<DetalleActividad />} />
          <Route path="/mis-actividades" element={<MisActividades />} />
          <Route path="/admin/nueva-actividad" element={<NuevaActividad />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;