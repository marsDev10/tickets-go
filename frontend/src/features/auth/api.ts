import { apiSlice } from '../../api/apiSlice'

import type { LoginPayload, SessionResponse, User } from './types'

export const authApi = apiSlice.injectEndpoints({
  endpoints: (build) => ({
    getSession: build.query<User | null, void>({
      query: () => ({
        url: '/auth/me',
        method: 'GET',
      }),
      providesTags: ['User'],
    }),
    login: build.mutation<SessionResponse, LoginPayload>({
      query: (body) => ({
        url: '/auth/login',
        method: 'POST',
        body,
      }),
      invalidatesTags: ['User'],
    }),
    logout: build.mutation<void, void>({
      query: () => ({
        url: '/auth/logout',
        method: 'POST',
      }),
      invalidatesTags: ['User'],
    }),
  }),
})

export const {
  useGetSessionQuery,
  useLoginMutation,
  useLogoutMutation,
} = authApi
