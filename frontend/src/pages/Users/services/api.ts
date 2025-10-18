import { apiSlice } from '../../../api/apiSlice'
import type { ApiTagType } from '../../../api/apiSlice'
import type { ApiResponse } from '../../../interfaces/Response'
import type { IGetUserResponse } from '../interfaces/Reponse'
import type { IUser, TCreateUser } from '../interfaces/User'

export const api = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUsersByOrganization: builder.query<IGetUserResponse, void>({
      query: () => ({
        url: '/users/organization/',
        method: 'GET',
      }),
      providesTags: ['Users' as ApiTagType],
    }),
    createUser: builder.mutation<ApiResponse<IUser>, TCreateUser>({
      query: (body) => ({
        url: '/users/',
        method: 'POST',
        body,
      }),
      invalidatesTags: ['Users' as ApiTagType],
    }),
  }),
})

export const {
  useGetUsersByOrganizationQuery,
  useCreateUserMutation,
} = api