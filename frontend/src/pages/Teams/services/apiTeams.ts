import { apiSlice } from '../../../api/apiSlice';
import type { IApiResponseTeamMember } from '../interfaces/apiTeams.interface';

export const apiTeams = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getTeamsByOrganization: builder.query<IApiResponseTeamMember, void>({
      query: () => '/teams/all',
    }),
    getTeamById: builder.query({
      query: (id) => `/teams/${id}`,
    }),
    createTeam: builder.mutation({
      query: (newTeam) => ({
        url: '/teams',
        method: 'POST',
        body: newTeam,
      }),
    }),
    updateTeam: builder.mutation({
      query: ({ id, ...updatedTeam }) => ({
        url: `/teams/${id}`,
        method: 'PUT',
        body: updatedTeam,
      }),
    }),
    deleteTeam: builder.mutation({
      query: (id) => ({
        url: `/teams/${id}`,
        method: 'DELETE',
      }),
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
