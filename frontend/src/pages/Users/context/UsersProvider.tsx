import { createContext, useContext, type ReactNode } from "react"
import { useUsers, type IUseUsers } from "../hooks/useUsers"



interface UsersContextValue {
    users: IUseUsers;
}

const UsersContext = createContext<UsersContextValue | null>(null)

interface UsersProviderProps {
  children: ReactNode
}

export const UsersProvider = ({ children }: UsersProviderProps) => {

  const users = useUsers();

  const value: UsersContextValue = {
    users,
  }

  return (
    <UsersContext.Provider value={value}>
      {children}
    </UsersContext.Provider>
  )
}

export const useUsersContext = (): UsersContextValue => {
  const context = useContext(UsersContext)

  if (!context) {
    throw new Error("useUsersContext must be used within a UsersProvider")
  }

  return context
}

export default UsersProvider