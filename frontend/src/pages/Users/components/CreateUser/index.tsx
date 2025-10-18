import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';

//Types
import { type TCreateUser } from '../../interfaces/User';
import { useUsersContext } from '../../context/UsersProvider';

const CreateUser = () => {

  const {
    setters: {
      setShowCreateUser
    },
    users: {
      handles: {
        handleCreateUser
      }
    }
  } = useUsersContext();

  const schema = yup
  .object()
  .shape({
    first_name: yup.string().required('Nombre es requerido'),
    last_name: yup.string().required('Apellido es requerido'),
    gender: yup.string().required('Género es requerido'),
    email: yup.string().email('Email inválido').required('Email es requerido'),
    password: yup.string().min(8, 'Mínimo 8 caracteres').required('Contraseña es requerida'),
    phone: yup.string().required('Teléfono es requerido'),
    role: yup.string().required('Rol es requerido'),
  })
  .required();

  const { register, handleSubmit, formState: { errors } } = useForm<TCreateUser>({
    resolver: yupResolver(schema)
  });

  const onSubmit = async (data: TCreateUser) => {

    const [error] = await handleCreateUser(data);

    if (error) {
      console.error("Error creating user:", error);
      return;
    } 
  }
  
  return (
    <div className="max-w-2xl mx-auto p-6 bg-primary text-white rounded-lg">
      <div className="mb-6">
        <p className="">Complete la información para crear un nuevo usuario</p>
      </div>

      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        {/* Nombre y Apellido */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label htmlFor="first_name" className="block text-sm font-medium mb-1">
              Nombre *
            </label>
            <input
              {...register('first_name')}
              type="text"
              id="first_name"
              className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.first_name ? 'border-red-500' : 'border-gray-300'
              }`}
              placeholder="Ingrese el nombre"
            />
            {errors.first_name && (
              <p className="mt-1 text-sm text-red-600">{errors.first_name.message}</p>
            )}
          </div>

          <div>
            <label htmlFor="last_name" className="block text-sm font-medium  mb-1">
              Apellido *
            </label>
            <input
              {...register('last_name')}
              type="text"
              id="last_name"
              className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.last_name ? 'border-red-500' : 'border-gray-300'
              }`}
              placeholder="Ingrese el apellido"
            />
            {errors.last_name && (
              <p className="mt-1 text-sm text-red-600">{errors.last_name.message}</p>
            )}
          </div>
        </div>

        {/* Email y Teléfono */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label htmlFor="email" className="block text-sm font-medium  mb-1">
              Email *
            </label>
            <input
              {...register('email')}
              type="email"
              id="email"
              className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.email ? 'border-red-500' : 'border-gray-300'
              }`}
              placeholder="usuario@ejemplo.com"
            />
            {errors.email && (
              <p className="mt-1 text-sm text-red-600">{errors.email.message}</p>
            )}
          </div>

          <div>
            <label htmlFor="phone" className="block text-sm font-medium  mb-1">
              Teléfono *
            </label>
            <input
              {...register('phone')}
              type="tel"
              id="phone"
              className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.phone ? 'border-red-500' : 'border-gray-300'
              }`}
              placeholder="+1 (555) 123-4567"
            />
            {errors.phone && (
              <p className="mt-1 text-sm text-red-600">{errors.phone.message}</p>
            )}
          </div>
        </div>

        {/* Género y Rol */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label htmlFor="gender" className="block text-sm font-medium mb-1">
              Género *
            </label>
            <select
              {...register('gender')}
              id="gender"
              className={`w-full px-3 py-2 border bg-white text-gray-900 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.gender ? 'border-red-500' : 'border-gray-300'
              }`}
            >
              <option value="">Seleccione género</option>
              <option value="male">Masculino</option>
              <option value="female">Femenino</option>
              <option value="other">Otro</option>
            </select>
            {errors.gender && (
              <p className="mt-1 text-sm text-red-600">{errors.gender.message}</p>
            )}
          </div>

          <div>
            <label htmlFor="role" className="block text-sm font-medium mb-1">
              Rol *
            </label>
            <select
              {...register('role')}
              id="role"
              className={`w-full px-3 py-2 border bg-white text-gray-900 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                errors.role ? 'border-red-500' : 'border-gray-300'
              }`}
            >
              <option value="">Seleccione rol</option>
              <option value="admin">Administrador</option>
              <option value="manager">Manager</option>
            </select>
            {errors.role && (
              <p className="mt-1 text-sm text-red-600">{errors.role.message}</p>
            )}
          </div>
        </div>

        {/* Contraseña */}
        <div>
          <label htmlFor="password" className="block text-sm font-medium  mb-1">
            Contraseña *
          </label>
          <input
            {...register('password')}
            type="password"
            id="password"
            className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
              errors.password ? 'border-red-500' : 'border-gray-300'
            }`}
            placeholder="Mínimo 8 caracteres"
          />
          {errors.password && (
            <p className="mt-1 text-sm text-red-600">{errors.password.message}</p>
          )}
        </div>

        {/* Botones */}
        <div className="flex gap-4 pt-4">
          <button
            type="submit"
            className="flex-1 bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition duration-200 font-medium"
          >
            Crear Usuario
          </button>
          <button
            type="button"
            onClick={() => setShowCreateUser(false)}
            className="flex-1 bg-gray-600 text-white py-2 px-4 rounded-md hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition duration-200 font-medium"
          >
            Cancelar
          </button>
        </div>
      </form>
    </div>
  )
}

export default CreateUser