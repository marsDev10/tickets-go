import { useMemo, useState, type ChangeEvent } from "react";
import { useCreateUserMutation, useGetUsersByOrganizationQuery, useUpdateUserMutation } from "../services/api";
import { type IGetUserResponse } from "../interfaces/Reponse";
import type { TCreateUser, TUpdateUser, TUserPartial } from "../interfaces/User";

export interface IUseUsers {
    state: {
        search: string;
        users: IGetUserResponse;
        filteredUsers: TUserPartial[];
    };
    setters: {
        setSearch: (search: string) => void;
    },
    loadings: {
        loadingGetUsers: boolean;
        loadingCreateUser: boolean;
        loadingUpdateUser: boolean;
    };
    handles: {
        handleCreateUser: (data: TCreateUser) => Promise<[unknown | null, any]>;
        handleUpdateUser: (data: TUpdateUser) => Promise<[unknown | null, any]>;
        handlerSearchUser: (e: ChangeEvent<HTMLInputElement>) => void;
    };
}

export const useUsers = (): IUseUsers => {

    const [search, setSearch] = useState("");

    const {
        data: { data } = { data: [] },
        isLoading: loadingGetUsers
    } = useGetUsersByOrganizationQuery();

    const [create, { isLoading: loadingCreateUser }] = useCreateUserMutation();
    const [update, { isLoading: loadingUpdateUser }] = useUpdateUserMutation();

    const users = useMemo<IGetUserResponse>(() => ({
        data: data ?? [],
        pagination: {
            limit: 0,
            page: 0,
            total: 0,
            total_pages: 0
        },
        status: false
    }), [data])

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

    const handlerSearchUser = (e:  ChangeEvent<HTMLInputElement>) => {
        setSearch(e.target.value);
    }

    const filteredUsers = useMemo(() => {
        if (!search) return users.data;

        const lowercasedSearch = search.toLowerCase();
        const filteredData = users.data.filter(user => 
            `${user.first_name} ${user.last_name}`.toLowerCase().includes(lowercasedSearch) ||
            user?.email?.toLowerCase().includes(lowercasedSearch)
        );

        return filteredData;
    }, [search, users]);

    console.log("Filtered Users:", filteredUsers);

  return {
    state: {
        search,
        users,
        filteredUsers,
    },
    setters: {
        setSearch,
    },
    loadings: {
        loadingGetUsers,
        loadingCreateUser,
        loadingUpdateUser,
    }, 
    handles: {
        handleCreateUser,
        handleUpdateUser,
        handlerSearchUser
    }
  }
}