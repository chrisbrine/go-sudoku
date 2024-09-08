import { createSlice, SerializedError } from "@reduxjs/toolkit";
import { EPage, GameData } from "../../types";

import * as gameActions from "../reducers/game";
import {
  getGame,
  newGame,
  setMove,
  setHint,
  setHintRemove,
  quitGame,
  doLogin,
  doLogout,
  doRegister,
  updateUsername,
  updatePassword,
  updateName,
  updateDifficulty,
} from "../actions/game";
import { IGameStateData } from "../types";

export const gameSlice = createSlice({
  name: "game",
  initialState: {
    game: {} as GameData,
    loading: true,
    loadingMove: false,
    error: undefined as unknown as SerializedError,
    pickHint: false,
    page: EPage.LOGIN,
    loggedIn: false,
  } as IGameStateData,
  reducers: {
    setPickHint: gameActions.setPickHint,
    setLoading: gameActions.setLoading,
    setPage: gameActions.setPage,
  },
  extraReducers: (builder) => {
    builder
      .addCase(getGame.pending, (state) => {
        state.loading = true;
      })
      .addCase(getGame.fulfilled, (state, action) => {
        state.game = action.payload;
        state.loggedIn = true;
        if (state.game.InGame) {
          state.page = EPage.GAME;
        } else {
          state.page = EPage.DASHBOARD;
        }
        state.loading = false;
      })
      .addCase(getGame.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(newGame.pending, (state) => {
        state.loading = true;
      })
      .addCase(newGame.fulfilled, (state, action) => {
        state.game = action.payload;
        state.page = EPage.GAME;
        state.pickHint = false;
        state.loading = false;
      })
      .addCase(newGame.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(setMove.pending, (state) => {
        state.loadingMove = true;
      })
      .addCase(setMove.fulfilled, (state, action) => {
        state.game = action.payload;
        state.loadingMove = false;
      })
      .addCase(setMove.rejected, (state, action) => {
        state.loadingMove = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(setHint.pending, (state) => {
        state.loadingMove = true;
      })
      .addCase(setHint.fulfilled, (state, action) => {
        state.game = action.payload;
        state.loadingMove = false;
      })
      .addCase(setHint.rejected, (state, action) => {
        state.loadingMove = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(setHintRemove.pending, (state) => {
        state.loadingMove = true;
      })
      .addCase(setHintRemove.fulfilled, (state, action) => {
        state.game = action.payload;
      })
      .addCase(setHintRemove.rejected, (state, action) => {
        state.loadingMove = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(quitGame.pending, (state) => {
        state.loading = true;
      })
      .addCase(quitGame.fulfilled, (state, action) => {
        state.game = action.payload;
        state.page = EPage.DASHBOARD;
        state.pickHint = false;
        state.loading = false;
      })
      .addCase(quitGame.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(doLogin.pending, (state) => {
        state.loading = true;
      })
      .addCase(doLogin.fulfilled, (state, action) => {
        state.loggedIn = true;
        state.loading = false;
        state.pickHint = false;
        state.game = action.payload;
        if (state.game.InGame) {
          state.page = EPage.GAME;
        } else {
          state.page = EPage.DASHBOARD;
        }
      })
      .addCase(doLogin.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(doLogout.pending, (state) => {
        state.loading = true;
      })
      .addCase(doLogout.fulfilled, (state) => {
        state.loading = false;
        state.loggedIn = false;
        state.pickHint = false;
        state.page = EPage.LOGIN;
      })
      .addCase(doLogout.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(doRegister.pending, (state) => {
        state.loading = true;
      })
      .addCase(doRegister.fulfilled, (state, action) => {
        state.loggedIn = true;
        state.loading = false;
        state.pickHint = false;
        state.game = action.payload;
        if (state.game.InGame) {
          state.page = EPage.GAME;
        } else {
          state.page = EPage.DASHBOARD;
        }
      })
      .addCase(doRegister.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(updateUsername.pending, (state) => {
        state.loading = true;
      })
      .addCase(updateUsername.fulfilled, (state, action) => {
        state.loading = false;
        state.game = action.payload;
      })
      .addCase(updateUsername.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(updatePassword.pending, (state) => {
        state.loading = true;
      })
      .addCase(updatePassword.fulfilled, (state, action) => {
        state.loading = false;
        state.game = action.payload;
      })
      .addCase(updatePassword.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(updateName.pending, (state) => {
        state.loading = true;
      })
      .addCase(updateName.fulfilled, (state, action) => {
        state.loading = false;
        state.game = action.payload;
      })
      .addCase(updateName.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      })
      .addCase(updateDifficulty.pending, (state) => {
        state.loading = true;
      })
      .addCase(updateDifficulty.fulfilled, (state, action) => {
        state.loading = false;
        state.game = action.payload;
      })
      .addCase(updateDifficulty.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error;
        console.error(action.error);
      });
  },
});

export const { setLoading, setPickHint, setPage } = gameSlice.actions;
export const gamesReducer = gameSlice.reducer;
