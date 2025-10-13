import { Search, Plus, MoreVertical } from 'lucide-react'

interface User {
  id: number
  name: string
  email: string
  role: 'Admin' | 'Agent' | 'User'
  team: string
  status: 'active' | 'inactive'
  lastLogin: string
  initials: string
}

const users: User[] = [
  {
    id: 1,
    name: 'Carlos Mendoza',
    email: 'carlos@helpdesk.com',
    role: 'Admin',
    team: 'Engineering',
    status: 'active',
    lastLogin: '2 hours ago',
    initials: 'CM'
  },
  {
    id: 2,
    name: 'Ana Garcia',
    email: 'ana@helpdesk.com',
    role: 'Agent',
    team: 'Support',
    status: 'active',
    lastLogin: '5 minutes ago',
    initials: 'AG'
  }
]

const Users = () => {
  return (
    <div className="flex flex-col gap-8 p-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold text-white">Users</h1>
          <p className="mt-1 text-slate-400">
            Manage your team members and their roles
          </p>
        </div>
        <button className="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-500">
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

      <div className="rounded-lg bg-slate-800 ring-1 ring-slate-700">
        <table className="w-full">
          <thead>
            <tr className="border-b border-slate-700">
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">User</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Role</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Team</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Status</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Last Login</th>
              <th className="w-20 px-6 py-4 text-right text-sm font-medium text-slate-400">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-700">
            {users.map((user) => (
              <tr key={user.id} className="group">
                <td className="px-6 py-4">
                  <div className="flex items-center gap-3">
                    <div className="flex h-10 w-10 items-center justify-center rounded-full bg-blue-600 text-sm font-medium text-white">
                      {user.initials}
                    </div>
                    <div>
                      <div className="font-medium text-slate-200">{user.name}</div>
                      <div className="text-sm text-slate-400">{user.email}</div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4">
                  <span className="rounded bg-slate-700 px-2.5 py-1 text-sm text-slate-200">
                    {user.role}
                  </span>
                </td>
                <td className="px-6 py-4 text-slate-200">{user.team}</td>
                <td className="px-6 py-4">
                  <span className="inline-flex items-center gap-1.5">
                    <span className={`h-2 w-2 rounded-full ${
                      user.status === 'active' ? 'bg-green-500' : 'bg-slate-500'
                    }`} />
                    <span className={user.status === 'active' ? 'text-green-500' : 'text-slate-500'}>
                      {user.status}
                    </span>
                  </span>
                </td>
                <td className="px-6 py-4 text-slate-400">{user.lastLogin}</td>
                <td className="px-6 py-4 text-right">
                  <button className="invisible rounded p-1 hover:bg-slate-700 group-hover:visible">
                    <MoreVertical className="h-5 w-5 text-slate-400" />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Users