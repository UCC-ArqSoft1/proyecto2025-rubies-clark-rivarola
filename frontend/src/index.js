import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css'; // esta ruta ahora es correcta
import App from './MyApp'; // corregida
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

reportWebVitals();