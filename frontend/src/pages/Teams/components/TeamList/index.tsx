import Loader from "../../../../components/Loaders/Loader";
import { getInitials } from "../../../../utils";
import { useTeamContext } from "../../context/TeamsProvider";

const teams = [
        {
            "id": 3,
            "name": "Team Sexo 2",
            "description": "Encargado del Sexo :D",
            "members": [
                {
                    "id": 3,
                    "first_name": "Erick",
                    "last_name": "Espinoza",
                    "email": "user@adryo.com.mx",
                    "role": "manager"
                },
                {
                    "id": 4,
                    "first_name": "User",
                    "last_name": "Example",
                    "email": "user2@adryo.com.mx",
                    "role": "manager"
                }
            ]
        },
        {
            "id": 4,
            "name": "Team Oraciones",
            "description": "Encargado de organizar la misa :D",
            "members": [
                {
                    "id": 4,
                    "first_name": "User",
                    "last_name": "Example",
                    "email": "user2@adryo.com.mx",
                    "role": "manager"
                },
                {
                    "id": 5,
                    "first_name": "User",
                    "last_name": "Example",
                    "email": "user3@adryo.com.mx",
                    "role": "manager"
                }
            ]
        }
    ];

const TeamList = () => {

    const {
        teams: {
            state: {
                teams,
            },
            loadings: {
                isLoadingTeams,
            }
        },
    } = useTeamContext();

    if(isLoadingTeams) {
        return <Loader/>
    }
    
  return (
    <section className="grid grid-cols-3 gap-5">
        {teams?.data?.map((team) => (
            <div key={team.id} className="bg-blue-500/10 border-blue-500/20 p-4 rounded-lg">
                <div className="flex justify-between">
                    <section className="flex flex-col mb-4">
                        <h2 className="text-xl font-semibold mb-2">{team.name}</h2>
                        <p className="text-slate-400 mb-4">{team.description}</p>
                    </section>
                    <section>
                        <h3 className="font-medium text-xs border border-white p-1 rounded">{team.members.length} Members</h3>
                    </section>
                </div>
                <section>
                    <div className="flex -space-x-2">
                        {team.members.map((member) => (
                            <span key={member.id} className="w-8 h-8 rounded-full  border-2 flex items-center justify-center text-xs font-medium bg-blue-500/10 border-blue-500/20 hover:border-blue-500/40 text-blue-400">
                                {getInitials(member.first_name + ' ' + member.last_name)}
                            </span>
                        ))}
                    </div>
                </section>
                <section className="mt-4">
                    <button className="w-full bg-transparent border border-white text-xs text-white px-4 py-2 rounded cursor-pointer hover:bg-blue-500/20">
                        Manage Team
                    </button>
                </section>
            </div>
        ))}
    </section>
  )
}

export default TeamList