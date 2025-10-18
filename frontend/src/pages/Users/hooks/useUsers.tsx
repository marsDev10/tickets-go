import { useMemo } from "react";
import { useCreateUserMutation, useGetUsersByOrganizationQuery, useUpdateUserMutation } from "../services/api";
import { type IGetUserResponse } from "../interfaces/Reponse";
import type { TCreateUser, TUpdateUser } from "../interfaces/User";

export interface IUseUsers {
    state: IGetUserResponse;
    loadings: {
        loadingGetUsers: boolean;
        loadingCreateUser: boolean;
        loadingUpdateUser: boolean;
    },
    handles: {
        handleCreateUser: (data: TCreateUser) => Promise<[unknown | null, any]>;
        handleUpdateUser: (data: TUpdateUser) => Promise<[unknown | null, any]>;
    }
}

export const useUsers = (): IUseUsers => {

    const {
        data: { data: users } = { data: [] },
        isLoading: loadingGetUsers
    } = useGetUsersByOrganizationQuery()

    const [create, { isLoading: loadingCreateUser }] = useCreateUserMutation();
    const [update, { isLoading: loadingUpdateUser }] = useUpdateUserMutation();

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

    const handleUpdateUser = async (data: TUpdateUser): Promise<[unknown | null, any]> => {
        try {
            
            const result = await update(data);

            console.log("User updated successfully:", result);

            return [null, result];

        } catch (error) {
            console.error("Error updating user:", error);
            return [error, null];
        }
    }

  return {
    state,
    loadings: {
        loadingGetUsers,
        loadingCreateUser,
        loadingUpdateUser,
    }, 
    handles: {
        handleCreateUser,
        handleUpdateUser,
    }
  }
}