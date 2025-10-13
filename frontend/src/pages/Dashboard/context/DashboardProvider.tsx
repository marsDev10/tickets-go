import { createContext, useContext, type ReactNode } from "react";
import { useSumaryDashboard, type ISumaryDashboard } from "../hooks/useSumaryDashboard";

interface DashboardContextValue {
  summary: ISumaryDashboard;
}

const DashboardContext = createContext<DashboardContextValue | null>(null);

interface DashboardProviderProps {
  children: ReactNode;
}

export const DashboardProvider = ({ children }: DashboardProviderProps) => {
  const summary = useSumaryDashboard();

  const value: DashboardContextValue = {
    summary,
  };

  return (
    <DashboardContext.Provider value={value}>
      {children}
    </DashboardContext.Provider>
  );
};

export const useDashboardContext = (): DashboardContextValue => {
  const context = useContext(DashboardContext);

  if (!context) {
    throw new Error("useDashboardContext must be used within a DashboardProvider");
  }

  return context;
};
