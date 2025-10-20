import { Plus, Search } from "lucide-react"
import { useUsersContext } from "../../context/UsersProvider"

const Header = () => {

  const {
    handles: {
      handleUserShowModal,
    }
  } = useUsersContext()

  return (
    <>
        <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold text-white">Users</h1>
          <p className="mt-1 text-slate-400">
            Manage your team members and their roles
          </p>
        </div>
        <button 
        onClick={() => handleUserShowModal(null)}
        className="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-500">
          <Plus size={20} />
          Add User
        </button>
      </div>

      <div className="flex items-center gap-4">
        <div className="relative flex-1">
          <Search className="absolute left-3 top-1/2 h-5 w-5 -translate-y-1/2 text-slate-400" />
          <input
            type="text"
            placeholder="Search by name or email..."
            className="w-full rounded-lg bg-slate-800 py-2 pl-10 pr-4 text-slate-200 placeholder-slate-400 outline-none ring-1 ring-slate-700 focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <select className="rounded-lg bg-slate-800 px-4 py-2 text-slate-200 outline-none ring-1 ring-slate-700">
          <option>All Roles</option>
          <option>Admin</option>
          <option>Agent</option>
          <option>User</option>
        </select>
      </div>
    </>
  )
}

export default Header