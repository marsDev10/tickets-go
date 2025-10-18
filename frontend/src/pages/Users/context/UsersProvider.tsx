import { createContext, useContext, useState, type ReactNode } from "react"
import { useUsers, type IUseUsers } from "../hooks/useUsers"



interface UsersContextValue {
    state: {
      showCreateUser: boolean;
      isEdit: boolean;
    },
    setters: {
      setShowCreateUser: (show: boolean) => void;
      setIsEdit: (isEdit: boolean) => void;
    };
    users: IUseUsers;
}

const UsersContext = createContext<UsersContextValue | null>(null)

interface UsersProviderProps {
  children: ReactNode
}

export const UsersProvider = ({ children }: UsersProviderProps) => {

  const [showCreateUser, setShowCreateUser] = useState(false);
  const [isEdit, setIsEdit] = useState(false);

  const users = useUsers();

  const value: UsersContextValue = {
    state: {
      showCreateUser,
      isEdit,
    },
    setters: {
      setShowCreateUser,
      setIsEdit,
    },
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