import { Suspense, lazy, type ReactNode } from 'react'
import { Navigate, Outlet, createBrowserRouter } from 'react-router'

import {
  RedirectIfAuthenticated,
  RequireAuth,
} from '../features/auth'
import { FullScreenSpinner } from '../components/feedback/FullScreenSpinner'
import { RouterErrorBoundary } from '../components/feedback/RouterErrorBoundary'
import { AppLayout } from '../layouts/AppLayout'
import { PublicLayout } from '../layouts/PublicLayout'

const DashboardPage = lazy(() => import('../pages/Dashboard'))
const Tickets = lazy(() => import('../pages/Tickets'))
const LoginPage = lazy(() => import('../pages/Login'))
const ForbiddenPage = lazy(() => import('../pages/Forbidden'))
const NotFoundPage = lazy(() => import('../pages/NotFound'))

const withSuspense = (element: ReactNode) => (
  <Suspense fallback={<FullScreenSpinner />}>{element}</Suspense>
)

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Outlet />,
    errorElement: <RouterErrorBoundary />,
    children: [
      { index: true, element: <Navigate to="/app/dashboard" replace /> },
      {
        element: (
          <RedirectIfAuthenticated pendingElement={<FullScreenSpinner />} />
        ),
        children: [
          {
            path: 'login',
            element: <PublicLayout />,
            children: [
              {
                index: true,
                element: withSuspense(<LoginPage />),
              },
            ],
          },
        ],
      },
      {
        element: <RequireAuth pendingElement={<FullScreenSpinner />} />,
        children: [
          {
            path: 'app',
            element: <AppLayout />,
            children: [
              {
                index: true,
                path: 'dashboard',
                element: withSuspense(<DashboardPage />),
              },
              {
                index: true,
                path: 'tickets',
                element: withSuspense(<Tickets />),
              },
              
            ],
          },
        ],
      },
      {
        path: 'forbidden',
        element: withSuspense(<ForbiddenPage />),
      },
      {
        path: '*',
        element: withSuspense(<NotFoundPage />),
      },
    ],
  },
])
