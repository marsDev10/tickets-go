import { useMemo } from "react";
import { useCreateUserMutation, useGetUsersByOrganizationQuery } from "../services/api";
import { type IGetUserResponse } from "../interfaces/Reponse";
import type { TCreateUser } from "../interfaces/User";

export interface IUseUsers {
    state: IGetUserResponse;
    loadings: {
        loadingGetUsers: boolean;
    },
    handles: {
        handleCreateUser: (data: TCreateUser) => Promise<[unknown | null, any]>;
    }
}

export const useUsers = (): IUseUsers => {

    const {
        data: { data: users } = { data: [] },
        isLoading: loadingGetUsers
    } = useGetUsersByOrganizationQuery()

    const [create, { isLoading: loadingCreateUser }] = useCreateUserMutation();

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

    const handleCreateUser = async (data: TCreateUser): Promise<[unknown | null, any]> => {
        try {
            
            const result = await create(data);

            console.log("User created successfully:", result);

            return [null, result];

        } catch (error) {
            console.error("Error creating user:", error);
            return [error, null];
        }
    }

  return {
    state,
    loadings: {
        loadingGetUsers
    }, 
    handles: {
        handleCreateUser,
    }
  }
}