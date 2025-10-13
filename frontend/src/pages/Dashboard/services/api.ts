import { apiSlice } from '../../../api/apiSlice'
import type { ApiTagType } from '../../../api/apiSlice'

interface DashboardStats {
  total_tickets: number
  open_tickets: number
  resolved_tickets: number
  closed_tickets: number
  tickets_by_priority: {
    high: number
    medium: number
    low: number
  }
  tickets_by_status: {
    open: number
    in_progress: number
    resolved: number
    closed: number
  }
  recent_tickets: {
    id: number
    ticket_number: string
    subject: string
    status: string
    created_at: string
  }[]
}

interface TeamPerformance {
  team_id: number
  team_name: string
  total_tickets: number
  resolved_tickets: number
  average_resolution_time: string
}

export const dashboardApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    // Obtener estadísticas generales
    getDashboardStats: builder.query<DashboardStats, void>({
      query: () => ({
        url: '/dashboard/stats',
        method: 'GET',
      }),
      providesTags: ['DashboardStats' as ApiTagType],
    }),

    // Obtener rendimiento por equipo
    getTeamPerformance: builder.query<TeamPerformance[], void>({
      query: () => ({
        url: '/dashboard/team-performance',
        method: 'GET',
      }),
      providesTags: ['TeamPerformance' as ApiTagType],
    }),

    // Obtener tickets por período
    getTicketsByPeriod: builder.query<any, { startDate: string; endDate: string }>({
      query: ({ startDate, endDate }) => ({
        url: '/dashboard/tickets-by-period',
        method: 'GET',
        params: { start_date: startDate, end_date: endDate },
      }),
      providesTags: ['TicketsByPeriod' as ApiTagType],
    }),
  }),
})

export const {
  useGetDashboardStatsQuery,
  useGetTeamPerformanceQuery,
  useGetTicketsByPeriodQuery,
} = dashboardApi