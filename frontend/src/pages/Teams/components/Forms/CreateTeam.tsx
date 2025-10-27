import { useForm } from 'react-hook-form';
import type { THandleTeam } from '../../interfaces/apiTeams.interface';
import { useTeamContext } from '../../context/TeamsProvider';

const CreateTeam = () => {

    const {
        setters: { setShowCreateTeam },
        teams: {
            handles: { handleTeam },
        }
    } = useTeamContext();

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<THandleTeam>();

    const onSubmit = async(data: THandleTeam) => {
        try {

            const res = await handleTeam(data);

            console.log({
                res,
            });

        } catch (error){
            console.log("Error: ", error);
        }
    }

  return (
    <div className="max-w-2xl mx-auto p-6 bg-primary text-white rounded-lg">
        <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-4">
                <div>
                    <label htmlFor="name" className="flex gap-2 text-sm font-medium">
                        Name
                        <span className="text-red-500 font-bold text-md">*</span>
                    </label>
                    <input
                    {...register('name')}
                    type="text"
                    id="name"
                    className={`w-full px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                        errors.name ? 'border-red-500' : 'border-gray-300'
                    }`}
                    placeholder="Ej: Frontend Team"
                    />
                    {errors.name && (
                    <p className="mt-1 text-sm text-red-600">{errors.name.message}</p>
                    )}
                </div>
                <div>
                    <label htmlFor="description" className="block text-sm font-medium  mb-1">
                        Description
                    </label>
                    <textarea
                    {...register('description')}
                    rows={4}
                    id="description"
                    className={`w-full px-3 resize-none py-2 border rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 ${
                        errors.description ? 'border-red-500' : 'border-gray-300'
                    }`}
                    placeholder="Ej: Front-end devepment team"
                    />
                    {errors.description && (
                    <p className="mt-1 text-sm text-red-600">{errors.description.message}</p>
                    )}
                </div>
            </div>
        <section
        className="mt-6 flex justify-end gap-4"
        >
            <button
            className="mt-4 bg-gray-600 hover:bg-gray-700 text-white font-semibold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
            onClick={() => setShowCreateTeam(false)}
            >
                Cancelar
            </button>
            <button
            type="submit"
            className="mt-4 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
            >
            Crear Equipo
            </button>
        </section>
        </form>
    </div>
  )
}
export default CreateTeam;