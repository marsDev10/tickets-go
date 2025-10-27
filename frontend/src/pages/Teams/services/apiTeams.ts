import { apiSlice } from '../../../api/apiSlice';
import type { IApiResponseCreateTeam, IApiResponseTeamMember, THandleTeam, } from '../interfaces/apiTeams.interface';

export const apiTeams = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getTeamsByOrganization: builder.query<IApiResponseTeamMember, void>({
      query: () => '/teams/all/',
      providesTags: ['Team'],
    }),
    getTeamById: builder.query({
      query: (id) => `/teams/${id}/`,
      providesTags: ['Team'],
    }),
    createTeam: builder.mutation<IApiResponseCreateTeam, THandleTeam>({
      query: (body) => ({
        url: '/teams/',
        method: 'POST',
        body,
      }),
      invalidatesTags: ['Team'],
    }),
    updateTeam: builder.mutation({
      query: ({ id, ...updatedTeam }) => ({
        url: `/teams/${id}/`,
        method: 'PUT',
        body: updatedTeam,
      }),
      invalidatesTags: ['Team'],
    }),
    deleteTeam: builder.mutation({
      query: (id) => ({
        url: `/teams/${id}/`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Team'],
    }),
  }),
});

export const {
  useGetTeamsByOrganizationQuery,
  useGetTeamByIdQuery,
  useCreateTeamMutation,
  useUpdateTeamMutation,
  useDeleteTeamMutation,
} = apiTeams;
