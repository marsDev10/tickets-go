import { Link } from 'react-router/dom'

const NotFound = () => (
  <div className="space-y-4 text-center text-slate-200">
    <h2 className="text-3xl font-semibold">404</h2>
    <p className="text-sm text-slate-400">
      No pudimos encontrar la página que buscabas.
    </p>
    <Link
      to="/"
      className="inline-flex items-center justify-center rounded-md bg-sky-500 px-4 py-2 text-sm font-medium text-white hover:bg-sky-400"
    >
      Volver al tablero
    </Link>
  </div>
)

export default NotFound
