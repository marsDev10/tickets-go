export interface IUser {
    ID:              number;
    CreatedAt:       Date;
    UpdatedAt:       Date;
    DeletedAt:       null;
    first_name:      string;
    last_name:       string;
    gender:          number;
    email:           string;
    phone:           string;
    password:        string;
    role:            string;
    is_active:       boolean;
    global_role:     string;
    organization_id: number;
}

export type IUserPartial = Partial<IUser>;
