import { Outlet } from 'react-router'

export const PublicLayout = () => (
  <div className="grid min-h-screen place-items-center bg-slate-950">
    <div className="w-full max-w-md rounded-xl border border-slate-800 bg-slate-900/80 p-8 shadow-lg backdrop-blur">
      <header className="mb-8 text-center">
        <h1 className="text-2xl font-semibold text-white">Helpdesk</h1>
        <p className="mt-2 text-sm text-slate-400">
          Ingresa con tus credenciales para continuar
        </p>
      </header>
      <Outlet />
    </div>
  </div>
)
