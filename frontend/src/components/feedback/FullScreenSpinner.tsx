export const FullScreenSpinner = () => (
  <div className="grid place-items-center bg-slate-950 text-slate-100">
    <div className="flex flex-col items-center gap-2">
      <span className="h-12 w-12 animate-spin rounded-full border-4 border-slate-700 border-t-sky-400" />
      <p className="text-sm font-medium uppercase tracking-wide text-slate-400">
        Cargando…
      </p>
    </div>
  </div>
)
