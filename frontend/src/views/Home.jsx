import React, { useEffect, useState } from 'react'
import api from '../api'
import './Home.css'

const Home = () => {
    const [allActivities, setAllActivities] = useState([])
    const [activities, setActivities] = useState([])
    const [error, setError] = useState(null)
    const [search, setSearch] = useState('')
    const [loadingId, setLoadingId] = useState(null)
    const userId = localStorage.getItem('user_id')

    // Carga inicial
    const loadActivities = async () => {
        try {
            const resp = await api.get('/activities')
            setAllActivities(resp.data)
            setActivities(resp.data)
        } catch (err) {
            console.error(err)
            setError('No se pudieron cargar las actividades')
        }
    }

    useEffect(() => {
        loadActivities()
    }, [])

    // Filtrado local por título, categoría, día u hora
    const handleSearch = e => {
        e.preventDefault()
        const term = search.trim().toLowerCase()
        if (!term) {
            setActivities(allActivities)
            return
        }
        const filtered = allActivities.filter(act => {
            const title = (act.title || act.titulo || '').toLowerCase()
            const category = (act.category_name || act.nombre_categoria || '').toLowerCase()
            const day = (act.day || act.dia_semana || '').toLowerCase()
            const time = (act.start_time || act.hora_inicio || '').toLowerCase()
            return (
                title.includes(term) ||
                category.includes(term) ||
                day.includes(term) ||
                time.includes(term)
            )
        })
        setActivities(filtered)
    }

    const handleSubscribe = async activityId => {
        if (!userId) return alert('Debes iniciar sesión primero')
        setLoadingId(activityId)
        try {
            const { data } = await api.get(`/activities/${activityId}`)
            const schedule = data.schedule
            if (!schedule || schedule.length === 0) {
                alert('No hay horarios disponibles')
                return
            }
            const horarioId = schedule[0].id
            await api.post('/inscripciones', {
                usuario_id: parseInt(userId, 10),
                horario_id: horarioId
            })
            alert('¡Inscripción exitosa!')
        } catch (err) {
            console.error(err)
            alert(err.response?.data?.error || 'Error al inscribirse')
        } finally {
            setLoadingId(null)
        }
    }

    if (error) return <p>{error}</p>

    return (
        <div className="home-container">
            <h1>Actividades Disponibles</h1>

            <form className="search-form" onSubmit={handleSearch}>
                <input
                    type="text"
                    placeholder="Buscar por título, categoría, día u horario..."
                    value={search}
                    onChange={e => setSearch(e.target.value)}
                />
                <button type="submit">Buscar</button>
            </form>

            <ul className="activities-list">
                {activities.map(act => {
                    // Construir el CSS de fondo solo si es URL absoluta
                    const bgImage = act.image_url && /^https?:\/\//.test(act.image_url)
                        ? `url(${act.image_url})`
                        : 'none'

                    return (
                        <li
                            key={act.id}
                            className="activity-card"
                            style={{ '--bg-image': bgImage }}
                        >
                            <div className="card-content">
                                <h2>{act.title || act.titulo}</h2>
                                <p><strong>Categoría:</strong> {act.category_name || act.nombre_categoria}</p>
                                <p><strong>Instructor:</strong> {act.instructor_name || act.nombre_instructor}</p>
                                <p>
                                    <strong>Próximo turno:</strong> {act.day || act.dia_semana} a las {act.start_time || act.hora_inicio}
                                </p>
                                <p><strong>Duración:</strong> {act.duration != null ? act.duration : act.duracion_min} minutos</p>
                                <p><strong>Cupo máximo:</strong> {act.capacity != null ? act.capacity : act.cupo_max} personas</p>
                                <button
                                    onClick={() => handleSubscribe(act.id)}
                                    disabled={loadingId === act.id}
                                >
                                    {loadingId === act.id ? 'Inscribiendo...' : 'Inscribirse'}
                                </button>
                            </div>
                        </li>
                    )
                })}
            </ul>
        </div>
    )
}

export default Home
