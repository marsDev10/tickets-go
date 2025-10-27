import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import type { RootState } from '../app/store'

const baseUrl = 'http://localhost:8080/api'

export const tagTypes = ['Dashboard', 'Ticket', 'User', 'Team', 'Settings', 'Auth'] as const;

export const apiSlice = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl,
    credentials: 'include',
    prepareHeaders: (headers, { getState }) => {

      // Obtener el token del estado de Redux
      const token = (getState() as RootState).auth.token

      if (token) {
        headers.set('Authorization',token);
      }

      return headers
    },
  }),
  tagTypes,
  endpoints: () => ({}),
})

export type ApiTagType = typeof tagTypes[number];
