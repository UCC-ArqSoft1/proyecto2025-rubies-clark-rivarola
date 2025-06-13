import React, { useEffect, useState } from 'react';

function Horarios() {
  const [horarios, setHorarios] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    fetch('http://localhost:8080/horarios', {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })
      .then(res => {
        if (!res.ok) throw new Error('Error al obtener horarios');
        return res.json();
      })
      .then(data => setHorarios(data))
      .catch(err => setError(err.message));
  }, []);

  return (
    <div>
      <h2>Horarios</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      {horarios.length === 0 ? (
        <p>No hay horarios disponibles</p>
      ) : (
        <ul>
          {horarios.map(h => (
            <li key={h.ID}>{h.Day} - {h.StartTime} - {h.DurationMin} min</li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Horarios;