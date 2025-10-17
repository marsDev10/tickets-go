import { createSlice, type PayloadAction } from '@reduxjs/toolkit';
import { api } from '../../pages/Users/services/api';

interface AuthState {
    user: null | { 
      id: number
      organization_id: number
      email: string
      first_name: string
      roles: string | string[]
    };
    token: null | string;
}

const initialState: AuthState = {
    user: localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')!) : null,
    token: localStorage.getItem('token') || null,
};

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        setCredentials(
            state,
            action: PayloadAction<{ user: AuthState['user']; token: string }>
        ) {
            state.user = action.payload.user;
            state.token = action.payload.token;
            // Guardar en localStorage
            localStorage.setItem('user', JSON.stringify(action.payload.user));
            localStorage.setItem('token', action.payload.token);
        },
        logout(state) {
            state.user = null;
            state.token = null;
            // Limpiar localStorage
            localStorage.removeItem('user');
            localStorage.removeItem('token');
            api.util.resetApiState();
            
        },
    },
});

export const { setCredentials, logout } = authSlice.actions;
export default authSlice.reducer;