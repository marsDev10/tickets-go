import { createContext, useContext } from 'react'

import type { User } from './types'
/* import type { FetchBaseQueryError } from '@reduxjs/toolkit/query'
import type { SerializedError } from '@reduxjs/toolkit'
 */
export interface AuthContextValue {
  user: User | null
  isAuthenticated: boolean
  isInitialising: boolean
  hasRole: (role: string | string[]) => boolean
}

export const AuthContext = createContext<AuthContextValue | null>(null)

export const useAuth = () => {
  const context = useContext(AuthContext)
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider')
  }

  return context
}
