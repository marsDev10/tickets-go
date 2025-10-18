import type { TUserPartial } from "./User";

export interface IGetUserResponse {
    data: TUserPartial[];
    pagination: {
        limit: number;
        page: number;
        total: number;
        total_pages: number;
    };
    status: boolean;
}