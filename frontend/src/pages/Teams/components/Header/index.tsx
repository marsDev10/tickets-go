import { Plus } from "lucide-react"
import { useTeamContext } from "../../context/TeamsProvider";

const Header = () => {

  const {
    setters: {
      setShowCreateTeam
    } 
  } = useTeamContext();

  return (
    <div  className="w-full">
        <section className="w-full flex justify-between items-center gap-4">
        <h1 className="text-2xl font-bold">Teams</h1>
        <button 
        onClick={() => setShowCreateTeam(true)}
        className="flex items-center bg-blue-500 text-white px-4 py-2 rounded cursor-pointer">
           <Plus className="inline-block mr-2 h-6 w-6"/>
            Create Team
        </button>
        </section>
        <section>
            <p className="text-sm text-slate-400">Manage your teams and their members here.</p>
        </section>
    </div>
  )
}

export default Header