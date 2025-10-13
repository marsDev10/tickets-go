import Loader from "../../../components/Loaders/Loader";
import { useDashboardContext } from "../context/DashboardProvider"

const Summary = () => {

    const {
        state: {
            OpenTickets,
            ClosedTickets,
            ResolvedTickets,
            TotalTickets
        },
        loadings: { 
            loadingSummary
        }
    } = useDashboardContext().summary;

    if (loadingSummary) {
        return (
            <Loader/>
        )
    }
    


  return (
    <section
    >
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
            <div className="rounded-lg border border-slate-800 bg-slate-900/60 p-4 text-center text-slate-300">
                <h3 className="text-lg font-medium">Total Tickets</h3>
                <p className="mt-2 text-2xl font-bold">{TotalTickets}</p>
                <p className="mt-1 text-sm text-slate-400">As of today</p>
            </div>
            <div className="rounded-lg border border-slate-800 bg-slate-900/60 p-4 text-center text-slate-300">
                <h3 className="text-lg font-medium">Open Tickets</h3>
                <p className="mt-2 text-2xl font-bold">{OpenTickets}</p>
                <p className="mt-1 text-sm text-slate-400">Currently open</p>
            </div>
            <div className="rounded-lg border border-slate-800 bg-slate-900/60 p-4 text-center text-slate-300">
                <h3 className="text-lg font-medium">Closed Tickets</h3>
                <p className="mt-2 text-2xl font-bold">{ClosedTickets}</p>
                <p className="mt-1 text-sm text-slate-400">Resolved this month</p>
            </div>
            <div className="rounded-lg border border-slate-800 bg-slate-900/60 p-4 text-center text-slate-300">
                <h3 className="text-lg font-medium">Resolded Tickets</h3>
                <p className="mt-2 text-2xl font-bold">{ResolvedTickets}</p>
                <p className="mt-1 text-sm text-slate-400">Awaiting response</p>
            </div>
        </div>
    </section>
  )
}

export default Summary