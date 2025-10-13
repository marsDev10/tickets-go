import { Navigate, Outlet } from 'react-router'
import type { ReactElement } from 'react'

import { useAuth } from '../context'

export interface RedirectIfAuthenticatedProps {
  redirectTo?: string
  pendingElement?: ReactElement | null
  children?: ReactElement
}

export const RedirectIfAuthenticated = ({
  redirectTo = '/app',
  pendingElement = null,
  children,
}: RedirectIfAuthenticatedProps) => {

  const { isAuthenticated, isInitialising } = useAuth()

  if (isInitialising) {
    return pendingElement
  }

  if (isAuthenticated) {
    return <Navigate to={redirectTo} replace />
  }

  if (children) {
    return children
  }

  return <Outlet />
}
