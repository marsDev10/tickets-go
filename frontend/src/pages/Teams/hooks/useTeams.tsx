import { useGetTeamsByOrganizationQuery } from "../services/apiTeams";

/* interface IUseTeamsReturn {
    state: {
        data: IApiResponseTeamMember | undefined;
    }
    loadings: {
        isLoading: boolean;
    }
} */

export const useTeams = () => {

    const { data: teams, error, isLoading: isLoadingTeams } = useGetTeamsByOrganizationQuery();

    console.log({ teams, error, isLoadingTeams });

  return {
    state: {
        teams,
    },
    loadings: {
        isLoadingTeams,
    }
  }
}
