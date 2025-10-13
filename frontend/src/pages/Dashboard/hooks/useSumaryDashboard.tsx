import { useMemo } from 'react';
import type { ISumaryDashboardData } from "../interfaces/response";
import { useGetDashboardSummaryQuery } from "../services/api";

export interface ISumaryDashboard {
    state: ISumaryDashboardData,
    loadings: {
        loadingSummary: boolean;
    }
}


export const useSumaryDashboard = (): ISumaryDashboard => {
    const { 
        data: { data: dashboardData } = { data: null }, 
        isLoading: loadingSummary 
    } = useGetDashboardSummaryQuery()

    const state = useMemo<ISumaryDashboardData>(() => ({
        TotalTickets: dashboardData?.TotalTickets ?? 0,
        OpenTickets: dashboardData?.OpenTickets ?? 0,
        ResolvedTickets: dashboardData?.ResolvedTickets ?? 0,
        ClosedTickets: dashboardData?.ClosedTickets ?? 0,
    }), [dashboardData])

    return {
        state,
        loadings: {
            loadingSummary
        }
    }
}
