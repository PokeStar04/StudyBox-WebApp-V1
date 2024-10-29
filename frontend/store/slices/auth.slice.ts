// // auth.slice.ts
// import { createSlice, PayloadAction } from '@reduxjs/toolkit';

// interface AuthState {
//   token: string | null;
//   isAuthenticated: boolean;
// }

// const initialState: AuthState = {
//   token: null,
//   isAuthenticated: false,
// };

// const authSlice = createSlice({
//   name: 'auth',
//   initialState,
//   reducers: {
//     login: (
//       state,
//       action: PayloadAction<{ token: string; isAuthenticated: boolean }>,
//     ) => {
//       state.token = action.payload.token; // Accès à la propriété token de l'objet payload
//       state.isAuthenticated = action.payload.isAuthenticated; // Accès à isAuthenticated
//     },
//     logout: (state) => {
//       state.token = null;
//       state.isAuthenticated = false;
//     },
//   },
// });

// export const { login, logout } = authSlice.actions;
// export default authSlice.reducer;

import { createSlice } from '@reduxjs/toolkit';

interface AuthState {
  token: string | null;
  isAuthenticated: boolean;
  role: string | null; // Ajoutez le rôle ici
}

const initialState: AuthState = {
  token: null,
  isAuthenticated: false,
  role: null,
};

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    login(state, action) {
      state.token = action.payload.token;
      state.isAuthenticated = action.payload.isAuthenticated;
      state.role = action.payload.role; // Stockez le rôle ici
    },
    logout(state) {
      state.token = null;
      state.isAuthenticated = false;
      state.role = null;
    },
    setRole(state, action) {
      state.role = action.payload; // Si le rôle est récupéré après la connexion
    },
  },
});

export const { login, logout, setRole } = authSlice.actions;
export default authSlice.reducer;
