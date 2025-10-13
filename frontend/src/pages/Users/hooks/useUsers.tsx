import { useMemo } from "react";
import { useGetUsersByOrganizationQuery } from "../services/api";
import { type IGetUserResponse } from "../interfaces/Reponse";

export interface IUseUsers {
    state: IGetUserResponse;
    loadings: {
        loadingGetUsers: boolean;
    }
}

export const useUsers = (): IUseUsers => {

    const {
        data: { data: users } = { data: [] },
        isLoading: loadingGetUsers
    } = useGetUsersByOrganizationQuery()

    const state = useMemo<IGetUserResponse>(() => ({
        data: users ?? [],
        pagination: {
            limit: 0,
            page: 0,
            total: 0,
            total_pages: 0
        },
        status: false
    }), [users])

  return {
    state,
    loadings: {
        loadingGetUsers
    }
  }
}