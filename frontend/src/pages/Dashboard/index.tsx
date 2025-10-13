import Summary from "./components/Summary"

const Dashboard = () => (
  <section className="space-y-6">
    <header>
      <h2 className="text-3xl font-semibold text-white">Dashboard</h2>
      <p className="mt-1 text-sm text-slate-400">
        Welcome back! Here's what's happening with your tickets today.
      </p>
    </header>
    <Summary/>
    <div className="grid gap-4 rounded-lg border border-slate-800 bg-slate-900/60 p-6 text-slate-300">
      <p>
        Esta vista es un prototipo. Aquí puedes renderizar dashboards,
        estadísticas, o accesos rápidos a distintas secciones de la
        plataforma.
      </p>
    </div>
  </section>
)

export default Dashboard
