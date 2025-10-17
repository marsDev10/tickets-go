import { useEffect, useState, type FormEvent } from 'react'
import { Link, useLocation, useNavigate } from 'react-router'
import type { FetchBaseQueryError } from '@reduxjs/toolkit/query'
import type { SerializedError } from '@reduxjs/toolkit'

import { useAuth, useLoginMutation } from '../../features/auth'
import { useDispatch } from 'react-redux'
import { setCredentials } from '../../features/auth/slice'
import { Eye, EyeClosed } from 'lucide-react'

const getErrorMessage = (
  apiError: FetchBaseQueryError | SerializedError | undefined,
) => {
  if (!apiError) return null
  if ('status' in apiError) {
    if (typeof apiError.data === 'string') return apiError.data
    if (
      apiError.data &&
      typeof apiError.data === 'object' &&
      'message' in apiError.data &&
      typeof apiError.data.message === 'string'
    ) {
      return apiError.data.message
    }
    return 'Servidor no disponible. Intenta más tarde.'
  }

  if ('message' in apiError && apiError.message) {
    return apiError.message
  }

  return 'Ocurrió un error inesperado.'
}

const Login = () => {

  const dispatch = useDispatch();

  const navigate = useNavigate()
  const location = useLocation()
  const { isAuthenticated } = useAuth()

  const [login, { isLoading, isError, error }] = useLoginMutation()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [remember, setRemember] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const redirectTo = (location.state as { from?: string } | null)?.from ?? '/app'

  const tocglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  }
  
  useEffect(() => {
    if (isAuthenticated) {
      navigate(redirectTo, { replace: true })
    }
  }, [isAuthenticated, navigate, redirectTo])

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    try {
      const response = await login({ email, password, remember }).unwrap();

      console.log()

      if(response.success) {
        dispatch(setCredentials({ user: response.data.user, token: response.data.token }));
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('user', JSON.stringify(response.data.user));
        navigate(redirectTo, { replace: true })
      }

    } catch (loginError) {
      console.error(loginError)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-6 text-left">
      <fieldset className="space-y-2">
        <label htmlFor="email" className="block text-sm font-medium text-slate-200">
          Correo electrónico
        </label>
        <input
          id="email"
          type="email"
          autoComplete="email"
          required
          value={email}
          onChange={(event) => setEmail(event.target.value)}
          className="w-full rounded-md border border-slate-700 bg-slate-950 px-3 py-2 text-slate-100 outline-none focus:border-sky-500 focus:ring-2 focus:ring-sky-500/40"
          placeholder="tu-correo@empresa.com"
        />
      </fieldset>

      <fieldset className="space-y-2">
        <label htmlFor="password" className="block text-sm font-medium text-slate-200">
          Contraseña
        </label>
        <section
        className="flex justify-between w-full rounded-md border border-slate-700 bg-slate-950 px-3 py-2 text-slate-100 outline-none focus:border-sky-500 focus:ring-2 focus:ring-sky-500/40"
        >
          <input
            id="password"
            type={showPassword ? 'text' : 'password'}
            autoComplete="current-password"
            required
            className='w-full bg-slate-950 focus:outline-none'
            value={password}
            onChange={(event) => setPassword(event.target.value)}
            placeholder="••••••"
          />
          <button
          type="button"
          onClick={tocglePasswordVisibility}
          className="flex items-center justify-center text-slate-400 hover:text-slate-200 focus:outline-none cursor-pointer"
          aria-label={showPassword ? 'Ocultar contraseña' : 'Mostrar contraseña'}
          >
            {
              showPassword ? (
                <EyeClosed/>
              ) : (
                <Eye />
              )
            }
          </button>
        </section>
      </fieldset>

      <label className="flex items-center gap-2 text-sm text-slate-300">
        <input
          type="checkbox"
          checked={remember}
          onChange={(event) => setRemember(event.target.checked)}
          className="h-4 w-4 rounded border-slate-700 bg-slate-900 accent-sky-500"
        />
        Mantener sesión iniciada
      </label>

      {isError && (
        <p className="rounded-md border border-transparent bg-red-500/10 px-3 py-2 text-sm text-red-400">
          Ocurrió un error al iniciar sesión. Verifica tus credenciales e intenta nuevamente.
          <br />
          <span className="text-xs text-red-500/70">{getErrorMessage(error)}</span>
        </p>
      )}

      <button
        type="submit"
        disabled={isLoading}
        className="w-full rounded-md bg-sky-500 px-4 py-2 text-sm font-semibold text-white shadow-lg transition hover:bg-sky-400 disabled:cursor-not-allowed disabled:opacity-60"
      >
        {isLoading ? 'Ingresando…' : 'Ingresar'}
      </button>

      <p className="text-center text-xs text-slate-500">
        ¿Olvidaste tu contraseña?{' '}
        <Link to="/support" className="font-medium text-sky-400 hover:text-sky-300">
          Contacta al equipo de soporte
        </Link>
      </p>
    </form>
  )
}

export default Login
