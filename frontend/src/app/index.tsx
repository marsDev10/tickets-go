import { RouterProvider } from 'react-router/dom'
import { AuthProvider } from '../features/auth'
import { router } from './router'

import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


const App = () => (
  <AuthProvider>
    <RouterProvider router={router} />
    <ToastContainer
    aria-label="Notifications"
    position="top-center"
    autoClose={3000}
    hideProgressBar={false}
    closeOnClick
    pauseOnHover
    />
  </AuthProvider>
)

export default App
