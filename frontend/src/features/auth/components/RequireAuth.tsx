import { Navigate, Outlet, useLocation } from 'react-router'
import type { ReactElement } from 'react'

import { useAuth } from '../context'

export interface RequireAuthProps {
  redirectTo?: string
  roles?: string[]
  pendingElement?: ReactElement | null
}

export const RequireAuth = ({
  redirectTo = '/login',
  roles,
  pendingElement = null,
}: RequireAuthProps) => {

  const location = useLocation()
  const { isAuthenticated, isInitialising, hasRole } = useAuth()

  if (isInitialising) {
    return pendingElement
  }

  if (!isAuthenticated) {
    return (
      <Navigate
        to={redirectTo}
        replace
        state={{ from: location.pathname + location.search }}
      />
    )
  }

  if (roles && roles.length > 0 && !hasRole(roles)) {
    return <Navigate to="/forbidden" replace />
  }

  return <Outlet />
}
