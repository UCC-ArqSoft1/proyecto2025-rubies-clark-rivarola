import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    const [actividades, setActividades] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        // Cambiá esta URL si tu backend corre en otro puerto o ruta
        fetch('http://localhost:8080/activities')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Error al obtener actividades');
                }
                return response.json();
            })
            .then(data => setActividades(data))
            .catch(err => setError(err.message));
    }, []);

    if (error) {
        return <div>Error: {error}</div>;
    }

    return (
        <div>
            <h1>Actividades Deportivas</h1>
            {actividades.length === 0 ? (
                <p>Cargando actividades...</p>
            ) : (
                <ul>
                    {actividades.map((actividad) => (
                        <li key={actividad.ID}>
                            <Link to={`/actividad/${actividad.ID}`}>
                                {actividad.Titulo} - {actividad.Horario} - {actividad.Instructor?.Nombre}
                            </Link>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default Home;