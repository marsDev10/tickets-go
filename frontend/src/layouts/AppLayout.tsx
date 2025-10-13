import { Outlet, useNavigate } from 'react-router'
import { useAuth } from '../features/auth'
import { useDispatch } from 'react-redux'
import { useState } from 'react'
import { 
  Ticket, 
  Menu, 
  X, 
  LayoutDashboard, 
  Users, 
  Inbox, 
  Search,
  Settings,
  Bell,
  Plus
} from 'lucide-react'

export const AppLayout = () => (
  <AuthenticatedShell />
)

const AuthenticatedShell = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const { user } = useAuth()
  const [isLoggingOut, setIsLoggingOut] = useState(false)
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false)

  const handleLogout = async () => {
    try {
      setIsLoggingOut(true)
      localStorage.clear()
      dispatch({ type: 'auth/logout' })
      navigate('/login', { replace: true })
    } catch (error) {
      console.error('No fue posible cerrar la sesión', error)
    } finally {
      setIsLoggingOut(false)
    }
  }

  const navItems = [
    { label: 'Dashboard', path: '/dashboard', icon: <LayoutDashboard size={20} /> },
    { label: 'Tickets', path: '/tickets', icon: <Inbox size={20} /> },
    { label: 'Users', path: '/users', icon: <Users size={20} /> },
    { label: 'Teams', path: '/teams', icon: <Users size={20} /> },
    { label: 'Settings', path: '/settings', icon: <Settings size={20} /> },
  ]

  return (
    <div className="flex h-screen bg-slate-950">
      {/* Sidebar */}
      <aside className="hidden w-64 flex-shrink-0 border-r border-slate-800 lg:flex lg:flex-col">
        <div className="flex h-16 items-center gap-2 border-b border-slate-800 px-4">
          <Ticket className="h-6 w-6 text-sky-500" />
          <span className="text-lg font-semibold tracking-wide text-white">Helpdesk</span>
        </div>
        
        <nav className="flex flex-1 flex-col gap-1 p-4">
          {navItems.map((item) => (
            <button
              key={item.path}
              onClick={() => navigate(item.path)}
              className="flex items-center gap-3 rounded-lg px-3 py-2 text-slate-300 transition hover:bg-slate-800 hover:text-white"
            >
              {item.icon}
              {item.label}
            </button>
          ))}
        </nav>
      </aside>

      {/* Main content */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Header */}
        <header className="flex h-16 items-center justify-between border-b border-slate-800 bg-slate-900/80 px-4 backdrop-blur">
          <div className="flex items-center gap-4">
            <button className="lg:hidden">
              <Menu size={24} />
            </button>
            <div className="flex items-center gap-2 rounded-lg bg-slate-800 px-3 py-1.5">
              <Search size={20} className="text-slate-400" />
              <input 
                type="text"
                placeholder="Search tickets, users, teams..."
                className="bg-transparent text-sm text-slate-300 outline-none placeholder:text-slate-500"
              />
            </div>
          </div>

          <div className="flex items-center gap-4">
            <button className="rounded-lg bg-sky-500 px-3 py-1.5 text-sm font-medium text-white shadow transition hover:bg-sky-400">
              <span className="flex items-center gap-2">
                <Plus size={18} />
                New Ticket
              </span>
            </button>
            <button className="relative">
              <Bell size={20} className="text-slate-400" />
              <span className="absolute -right-1 -top-1 flex h-4 w-4 items-center justify-center rounded-full bg-red-500 text-xs font-medium text-white">
                2
              </span>
            </button>
            <img
              src={`https://ui-avatars.com/api/?name=${user?.first_name}&background=random`}
              alt="Profile"
              className="h-8 w-8 rounded-full"
            />
          </div>
        </header>

        {/* Page content */}
        <main className="flex-1 overflow-auto bg-slate-900 p-6">
          <Outlet />
        </main>
      </div>
    </div>
  )
}
