import type { IUserPartial } from "./User";

export interface IGetUserResponse {
    data: IUserPartial[];
    pagination: {
        limit: number;
        page: number;
        total: number;
        total_pages: number;
    };
    status: boolean;
}