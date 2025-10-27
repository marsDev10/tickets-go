import { isFulfilled, isPending, isRejectedWithValue, type Middleware, type SerializedError } from '@reduxjs/toolkit'
import type { FetchBaseQueryError,  } from '@reduxjs/toolkit/query'
import { toast } from 'react-toastify'
import { store } from './store'
import { logout } from '../features/auth/slice'

const isFetchBaseQueryError = (error: unknown): error is FetchBaseQueryError => {
  return typeof error === 'object' && error !== null && 'status' in error
}

const isSerializedError = (error: unknown): error is SerializedError => {
  if (typeof error !== 'object' || error === null) {
    return false
  }

  const maybeSerialized = error as Partial<SerializedError>
  return (
    'message' in maybeSerialized ||
    'code' in maybeSerialized ||
    'name' in maybeSerialized ||
    'stack' in maybeSerialized
  )
}

/**
 * Middleware que reacciona a los cambios de estado de cualquier request
 * que pase por RTK Query (queries o mutations). Puedes despachar acciones,
 * mostrar notificaciones o disparar side effects en cada bloque.
 */
export const requestStatusMiddleware: Middleware = () => (next) => (action) => {

  //console.log("Action", action);

  if (isPending(action) && action.meta?.arg && typeof action.meta.arg === 'object' && action.meta.arg !== null && 'type' in action.meta.arg && 'endpointName' in action.meta.arg) {
    /* console.info(
      `Pending ${(action.meta.arg as any).type} for "${(action.meta.arg as any).endpointName}"`,
    ) */
  }

  if (isFulfilled(action) && action.meta?.arg && typeof action.meta.arg === 'object' && action.meta.arg !== null && 'type' in action.meta.arg && 'endpointName' in action.meta.arg) {
    /* console.info(
      `Success ${(action.meta.arg as any).type} for "${(action.meta.arg as any).endpointName}"`,
    ) */

    //console.log("Meta Success", action.meta);

    const method = (action.meta as any)?.baseQueryMeta?.request?.method as string | undefined;
    const message = (action.payload as any)?.message as string | undefined;

    //console.log("Message", message);

  if (method && typeof method === 'string') {
      //console.log("Method Success", method);
      
      // Solo mostrar toast para ciertos métodos
      if (['POST', 'PUT', 'DELETE', 'PATCH'].includes(method.toUpperCase())) {
        toast.success(message, {
          position: "bottom-right",
          autoClose: 3000,
          hideProgressBar: false,
          closeOnClick: false,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "colored",
        });
      }
    }
  }

  if (isRejectedWithValue(action)) {
    const errorPayload = action.payload

    console.log({ errorPayload });

    if (isFetchBaseQueryError(errorPayload) && typeof errorPayload.status === 'number') {
      switch (errorPayload.status) {
        case 400:
            toast.error((errorPayload.data as any)?.message || 'Error al realizar la acción', {
              autoClose: 5000,
              hideProgressBar: false,
              closeOnClick: false,
              pauseOnHover: true,
              draggable: true,
              progress: undefined,
              theme: "colored",
            }); 
            break;
        case 401:
        
          toast.error('Sesion expirada, redirigiendo al login...', {
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: false,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "colored",
          });
          
          store.dispatch(logout());
          window.location.href = '/login'
          break
        case 403: 
          toast.error('Credenciales inválidas', {
              autoClose: 5000,
              hideProgressBar: false,
              closeOnClick: false,
              pauseOnHover: true,
              draggable: true,
              progress: undefined,
              theme: "colored",
            }); 
            break;
        case 500:
         toast.error('!Ops, ocurrió un error interno', {
            autoClose: 5000,
            hideProgressBar: false,
            closeOnClick: false,
            pauseOnHover: true,
            draggable: true,
            progress: undefined,
            theme: "colored",
          });
          break
        default:
          console.warn(`Error ${errorPayload.status} en la solicitud.`)
      }
    } else if (isSerializedError(errorPayload)) {
      console.error('Ocurrio un error en la solicitud.', errorPayload)
    } else {
      console.error('Ocurrio un error no controlado en la solicitud.', errorPayload)
    }
  }

  return next(action)
}
