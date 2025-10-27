
export interface IMember {
    id: number;
    first_name: string;
    last_name: string;
    email: string;
    role: string;
}

export interface ITeam {
    id: number;
    name: string;
    description: string;
    members: IMember[];
}

export interface IApiResponseTeamMember {
    data: ITeam[];
    message: string;
    success: boolean;
}

export type TMemberPartial = Omit<IMember, 'id'>;