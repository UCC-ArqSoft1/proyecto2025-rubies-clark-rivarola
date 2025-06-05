import React from 'react';
import { useParams } from 'react-router-dom';

const DetalleActividad = () => {
    const { id } = useParams();

    return (
        <div>
            <h1>Detalle de Actividad</h1>
            <p>Mostrando detalles para la actividad con ID: {id}</p>
        </div>
    );
};

export default DetalleActividad;
