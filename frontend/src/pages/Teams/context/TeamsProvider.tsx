import { createContext, useContext, type ReactNode } from "react";
import { useTeams } from "../hooks/useTeams";

interface ITeamsContext {
    teams: ReturnType<typeof useTeams>;
}

export const TeamsContext = createContext<ITeamsContext | null>(null);

interface Props {
  children: ReactNode;
}

const TeamsProvider = ({ children }: Props) => {

    const teams = useTeams();

  return (
    <TeamsContext.Provider value={{ teams }}>
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