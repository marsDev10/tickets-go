import type { ISumaryDashboard } from './../interfaces/response';
import { apiSlice } from '../../../api/apiSlice'
import type { ApiTagType } from '../../../api/apiSlice'



export const dashboardApi = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getDashboardSummary: builder.query<ISumaryDashboard, void>({
      query: () => ({
        url: '/dashboard/summary',
        method: 'GET',
      }),
      providesTags: ['DashboardStats' as ApiTagType],
    }),
  }),
})

export const {
  useGetDashboardSummaryQuery,
} = dashboardApi