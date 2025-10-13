import { Outlet, useNavigate } from 'react-router'

import { useAuth } from '../features/auth'
import { useDispatch } from 'react-redux'
import { useState } from 'react'

export const AppLayout = () => (
  <AuthenticatedShell />
)

const AuthenticatedShell = () => {

  const dispatch = useDispatch();

  const navigate = useNavigate()
  const { user } = useAuth()

  const [isLoggingOut, setIsLoggingOut] = useState(false);

  

  const handleLogout = async () => {
    try {

      setIsLoggingOut(true);

      localStorage.clear();
      dispatch({ type: 'auth/logout' });
      navigate('/login', { replace: true })
    
    } catch (error) {
      console.error('No fue posible cerrar la sesión', error)
    } finally {
      setIsLoggingOut(false);
    }
  }

  return (
    <div className="min-h-screen bg-slate-950 text-slate-100">
      <header className="border-b border-slate-800 bg-slate-900/80 backdrop-blur">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-4 py-4">
          <p className="text-lg font-semibold tracking-wide">Helpdesk</p>
          <nav aria-label="Acciones de usuario" className="flex items-center gap-3">
            <span className="text-sm text-slate-300">
              {user?.first_name ?? user?.email ?? 'Usuario'}
            </span>
            <button
              type="button"
              onClick={handleLogout}
              disabled={isLoggingOut}
              className="rounded-md bg-sky-500 px-3 py-1.5 text-sm font-medium text-white shadow transition hover:bg-sky-400 disabled:cursor-not-allowed disabled:opacity-60"
            >
              {isLoggingOut ? 'Cerrando…' : 'Cerrar sesión'}
            </button>
          </nav>
        </div>
      </header>
      <main className="mx-auto max-w-6xl px-4 py-8">
        <Outlet />
      </main>
    </div>
  )
}
