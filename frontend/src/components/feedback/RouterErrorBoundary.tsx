import { isRouteErrorResponse, Link, useRouteError } from 'react-router'

export const RouterErrorBoundary = () => {
  const error = useRouteError()

  if (isRouteErrorResponse(error)) {
    return (
      <div className="grid min-h-screen place-items-center bg-slate-950 text-slate-100">
        <div className="space-y-4 text-center">
          <h1 className="text-3xl font-semibold">Error {error.status}</h1>
          <p className="text-sm text-slate-400">{error.statusText}</p>
        </div>
      </div>
    )
  }

  return (
    <div className="grid min-h-screen place-items-center bg-slate-950 text-slate-100">
      <div className="space-y-4 text-center">
        <h1 className="text-3xl font-semibold">Algo salió mal</h1>
        <p className="text-sm text-slate-400">
          Intenta refrescar la página o regresar al inicio.
        </p>
        <Link
          to="/"
          className="inline-flex items-center justify-center rounded-md bg-sky-500 px-4 py-2 text-sm font-medium text-white hover:bg-sky-400"
        >
          Volver al inicio
        </Link>
      </div>
    </div>
  )
}
