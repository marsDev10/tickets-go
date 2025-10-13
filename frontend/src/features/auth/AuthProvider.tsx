import { useMemo, useEffect, type PropsWithChildren } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { AuthContext, type AuthContextValue } from './context'
import { setCredentials } from './slice'

interface RootState {
  auth: {

    user: null | { 
      id: number
      organization_id: number
      email: string
      first_name: string
      roles: string | string[]
    }
    token: string | null,
  }

}

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const dispatch = useDispatch()
  const data = useSelector((state: RootState) => state.auth);
  
  // Cargar datos del localStorage al iniciar
  useEffect(() => {
    const storedUser = localStorage.getItem('user')
    const storedToken = localStorage.getItem('token')
    
    if (storedUser && storedToken) {
        const user = JSON.parse(storedUser)
        dispatch(setCredentials({ user, token: storedToken }))
    }
  }, [dispatch])

  const defaultValue: AuthContextValue = {
    user: null,
    isAuthenticated: false,
    isInitialising: false,
    hasRole: () => false,
  }
  
  const value = useMemo<AuthContextValue>(
    () => {

      if (!data?.user || !data?.token) return defaultValue

      return {
        user: data?.user,
        isAuthenticated: true,
        isInitialising: false,
        refresh: async () => {
          // Implementar lógica de refresh si es necesario
          return null
        },
        hasRole: (roleOrRoles) => {
          const roles = Array.isArray(roleOrRoles)
            ? roleOrRoles
            : [roleOrRoles]
          const assignedRoles = data.user?.roles ?? ""
          return roles.some((role) => assignedRoles.includes(role))
        },
      }
    },
    [data?.user, data?.token],
  )

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}
