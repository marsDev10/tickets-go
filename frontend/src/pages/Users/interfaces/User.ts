export interface IUser {
    ID:              number;
    CreatedAt:       Date;
    UpdatedAt:       Date;
    DeletedAt:       null;
    first_name:      string;
    last_name:       string;
    gender:          string;
    email:           string;
    phone:           string;
    password:        string;
    role:            string;
    is_active:       boolean;
    global_role:     string;
    organization_id: number;
}


export type TUserPartial = Partial<IUser>;

export type TCreateUser = Pick<IUser, 'first_name' | 'last_name' | 'gender' | 'email' | 'password' | 'phone' | 'role'>;
