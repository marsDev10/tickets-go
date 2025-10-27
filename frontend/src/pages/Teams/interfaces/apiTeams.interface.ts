
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
    members: IMember[] | null;
}

export interface IApiResponseTeamMember {
    data: ITeam[];
    message: string;
    success: boolean;
}



export type TMemberPartial = Omit<IMember, 'id'>;


export type THandleTeam = Partial<ITeam>;

export interface IApiResponseCreateTeam {
    data: {
        id: number;
        name: string;
        created_at: string;
        updated_at: string;
        organization_id: number;
    };
    message: string;
    success: boolean;   
}


