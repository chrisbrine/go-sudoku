import { createAsyncThunk } from "@reduxjs/toolkit";
import { API } from "../../api";

import { setLoading, setPickHint, setPage } from "../slices/game";
import { IRowColNum } from "../../types";

export { setLoading, setPickHint, setPage };

export const getGame = createAsyncThunk("game/getGame", async () => {
  return API.get();
});

export const newGame = createAsyncThunk("game/newGame", async () => {
  return API.new();
});

export const setMove = createAsyncThunk(
  "game/move",
  async ({ row, col, num }: IRowColNum) => {
    return API.move({ row, col, num });
  }
);

export const setHint = createAsyncThunk(
  "game/hint",
  async ({ row, col, num }: IRowColNum) => {
    return API.hint({ row, col, num });
  }
);

export const setHintRemove = createAsyncThunk(
  "game/hintRemove",
  async ({ row, col, num }: IRowColNum) => {
    return API.hintRemove({ row, col, num });
  }
);

export const quitGame = createAsyncThunk(
  "game/quitGame",
  async (_, { dispatch }) => {
    const gameData = await API.quit();
    dispatch(updateLeaderboard());
    return gameData;
  }
);

export const doLogin = createAsyncThunk(
  "game/login",
  async (
    { username, password }: { username: string; password: string },
    { dispatch }
  ) => {
    const gameData = await API.login(username, password);
    dispatch(updateLeaderboard());
    return gameData;
  }
);

export const doLogout = createAsyncThunk("game/logout", async () => {
  return API.logout();
});

export const doRegister = createAsyncThunk(
  "game/register",
  async ({
    username,
    name,
    password,
  }: {
    username: string;
    name: string;
    password: string;
  }) => {
    return API.register(username, password, name);
  }
);

export const updateUsername = createAsyncThunk(
  "game/updateUsername",
  async (username: string) => {
    return API.updateUserName(username);
  }
);

export const updatePassword = createAsyncThunk(
  "game/updatePassword",
  async ({
    oldPassword,
    password,
  }: {
    oldPassword: string;
    password: string;
  }) => {
    return API.updatePassword(oldPassword, password);
  }
);

export const updateName = createAsyncThunk(
  "game/updateName",
  async (name: string) => {
    return API.updateName(name);
  }
);

export const updateDifficulty = createAsyncThunk(
  "game/updateDifficulty",
  async (difficulty: number) => {
    return API.updateDifficulty(difficulty);
  }
);

export const updateLeaderboard = createAsyncThunk(
  "game/updateLeaderboard",
  async () => {
    return API.leaderboard();
  }
);
