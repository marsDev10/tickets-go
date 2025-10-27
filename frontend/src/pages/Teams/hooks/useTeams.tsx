
// Interfaces
import type { THandleTeam } from "../interfaces/apiTeams.interface";
import type { IResponseHandleTeam } from "../interfaces/useTeam.interface";

// Apis
import { useCreateTeamMutation, useGetTeamsByOrganizationQuery } from "../services/apiTeams";

export const useTeams = () => {

    const { data: teams, isLoading: isLoadingTeams } = useGetTeamsByOrganizationQuery();

    const [create] = useCreateTeamMutation();

    const handleTeam = async (team: THandleTeam): Promise<IResponseHandleTeam> => {
        try {

            const response = await create(team).unwrap();

            return [null, response];

        } catch(error){
            console.error("Error creating team:", error);
            return [true, null];
        }
    }

  return {
    state: {
        teams,
    },
    loadings: {
        isLoadingTeams,
    },
    handles: {
        handleTeam,
    },
  }
}
