import { Link } from 'react-router'

const Forbidden = () => (
  <div className="space-y-4 text-center text-slate-200">
    <h2 className="text-2xl font-semibold text-red-400">Acceso restringido</h2>
    <p className="text-sm text-slate-400">
      Tu cuenta no tiene permisos suficientes para acceder a este recurso.
    </p>
    <Link
      to="/"
      className="inline-flex items-center justify-center rounded-md bg-sky-500 px-4 py-2 text-sm font-medium text-white hover:bg-sky-400"
    >
      Volver al inicio
    </Link>
  </div>
)

export default Forbidden
