import { createContext, useContext, useState, type ReactNode } from "react"
import { useUsers, type IUseUsers } from "../hooks/useUsers"
import type { TUserPartial } from "../interfaces/User";

interface UsersContextValue {
    state: {
      showCreateUser: boolean;
      currentUser: TUserPartial | null;
    },
    setters: {
      setShowCreateUser: (show: boolean) => void;
      setCurrentUser: (user: TUserPartial | null) => void;
    };
    handles: {
      handleUserShowModal: (user: TUserPartial | null) => void;
    };
    users: IUseUsers;
}

const UsersContext = createContext<UsersContextValue | null>(null)

interface UsersProviderProps {
  children: ReactNode
}

export const UsersProvider = ({ children }: UsersProviderProps) => {

  const [showCreateUser, setShowCreateUser] = useState(false);
  const [currentUser, setCurrentUser] = useState<TUserPartial | null>(null);

  const handleUserShowModal = (user: TUserPartial | null) => {
    setCurrentUser(user);
    setShowCreateUser(!showCreateUser);
  }

  const users = useUsers();

  const value: UsersContextValue = {
    state: {
      showCreateUser,
      currentUser,
    },
    setters: {
      setShowCreateUser,
      setCurrentUser,
    },
    handles: {
      handleUserShowModal,
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