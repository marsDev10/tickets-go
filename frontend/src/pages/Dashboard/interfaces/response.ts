
export interface ISumaryDashboardData {
    TotalTickets: number;
    OpenTickets: number;
    ResolvedTickets: number;
    ClosedTickets: number;
}

export interface ISumaryDashboard{
  data: ISumaryDashboardData;
  message: string;
  success: boolean;
}