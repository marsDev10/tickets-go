import { createContext, useContext, useState, type ReactNode } from "react";
import { useTeams } from "../hooks/useTeams";

interface ITeamsContext {
  state: {
    showCreateTeam: boolean;
  },
  setters: {
    setShowCreateTeam: (value: boolean) => void;
  }, 
  teams: ReturnType<typeof useTeams>;

}

export const TeamsContext = createContext<ITeamsContext | null>(null);

interface Props {
  children: ReactNode;
}

const TeamsProvider = ({ children }: Props) => {

    const [showCreateTeam, setShowCreateTeam] = useState(false);

    const teams = useTeams();

  return (
    <TeamsContext.Provider value={{ 
      state: {
         showCreateTeam
      }, 
      setters: {
        setShowCreateTeam
      },
      teams,
    }}>
      {children}
    </TeamsContext.Provider>
  )
}

export default TeamsProvider

export const useTeamContext = () => {
    const context = useContext(TeamsContext);
    if (!context) {
        throw new Error("useTeamContext must be used within a TeamsProvider");
    }
    return context;
}