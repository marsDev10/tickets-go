import { apiSlice } from '../../../api/apiSlice'
import type { ApiResponse } from '../../../interfaces/Response'
import type { IGetUserResponse } from '../interfaces/Reponse'
import type { IUser, TCreateUser, TUpdateUser } from '../interfaces/User'

export const api = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUsersByOrganization: builder.query<IGetUserResponse, void>({
      query: () => ({
        url: '/users/organization/',
        method: 'GET',
      }),
      providesTags: ['User'],
    }),
    createUser: builder.mutation<ApiResponse<IUser>, TCreateUser>({
      query: (body) => ({
        url: '/users/',
        method: 'POST',
        body,
      }),
      invalidatesTags: ['User'],
    }),
    updateUser: builder.mutation<ApiResponse<IUser>, TUpdateUser>({
      query: (body) => ({
        url: '/users/',
        method: 'PUT',
        body,
      }),
      invalidatesTags: ['User'],
    }),
  }),
})

export const {
  useGetUsersByOrganizationQuery,
  useCreateUserMutation,
  useUpdateUserMutation,
} = api