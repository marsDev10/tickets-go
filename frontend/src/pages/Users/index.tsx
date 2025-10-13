import { MoreVertical } from 'lucide-react'

// Components
import Header from './components/Header'

// Interfaces
import type { IUserPartial } from './interfaces/User'
import { useUsersContext } from './context/UsersProvider'
import Loader from '../../components/Loaders/Loader'

const users: IUserPartial[] = [
  {
    ID: 1,
    first_name: 'Carlos Mendoza',
    email: 'carlos@helpdesk.com',
    role: 'Admin',
    is_active: true,
    UpdatedAt: new Date(),

  },
  {
    ID: 2,
    first_name: 'Ana Garcia',
    email: 'ana@helpdesk.com',
    role: 'Agent',
    is_active: true,
    UpdatedAt: new Date(),
  }
]

const Users = () => {

  const {
    state: { data: users },
    loadings: { loadingGetUsers },
  } = useUsersContext().users;

  const getInitials = (name: string) => {
    return name.split(' ').map(n => n[0]).join('').toUpperCase()
  }

  if (loadingGetUsers) {
    return (
      <Loader/>
    )
  }

  return (
    <div className="flex flex-col gap-8 p-8">
      <Header/>

      <div className="rounded-lg bg-slate-800 ring-1 ring-slate-700">
        <table className="w-full">
          <thead>
            <tr className="border-b border-slate-700">
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">User</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Role</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Status</th>
              <th className="px-6 py-4 text-left text-sm font-medium text-slate-400">Last Updated</th>
              <th className="w-20 px-6 py-4 text-right text-sm font-medium text-slate-400">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-700">
            {users.map((user) => (
              <tr key={user.ID} className="group">
                <td className="px-6 py-4">
                  <div className="flex items-center gap-3">
                    <div className="flex h-10 w-10 items-center justify-center rounded-full bg-blue-600 text-sm font-medium text-white">
                      {getInitials(user.first_name || '')}
                    </div>
                    <div>
                      <div className="font-medium text-slate-200">{user.first_name}</div>
                      <div className="text-sm text-slate-400">{user.email}</div>
                    </div>
                  </div>
                </td>
                <td className="px-6 py-4">
                  <span className="rounded bg-slate-700 px-2.5 py-1 text-sm text-slate-200">
                    {user.role}
                  </span>
                </td>
                <td className="px-6 py-4">
                  <span className="inline-flex items-center gap-1.5">
                    <span className={`h-2 w-2 rounded-full ${
                      user.is_active ? 'bg-green-500' : 'bg-slate-500'
                    }`} />
                    <span className={user.is_active ? 'text-green-500' : 'text-slate-500'}>
                      {user.is_active ? 'Active' : 'Inactive'}
                    </span>
                  </span>
                </td>
                <td className="px-6 py-4 text-slate-400">
                  {user.UpdatedAt ? new Date(user.UpdatedAt).toLocaleDateString() : 'N/A'}
                </td>
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