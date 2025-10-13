import { apiSlice } from '../../../api/apiSlice'
import type { ApiTagType } from '../../../api/apiSlice'
import type { IGetUserResponse } from '../interfaces/Reponse'

export const api = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getUsersByOrganization: builder.query<IGetUserResponse, void>({
      query: () => ({
        url: '/users/organization',
        method: 'GET',
      }),
      providesTags: ['Users' as ApiTagType],
    }),
  }),
})

export const {
  useGetUsersByOrganizationQuery,
} = api