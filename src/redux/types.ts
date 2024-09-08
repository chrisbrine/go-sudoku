import type { Action, SerializedError, ThunkAction } from "@reduxjs/toolkit";
import { store } from "./store";
import { EPage, GameData, ILeaderboard } from "../types";

export type AppStore = typeof store;
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = AppStore["dispatch"];
export type AppThunk<ThunkReturnType = void> = ThunkAction<
  ThunkReturnType,
  RootState,
  unknown,
  Action
>;

export type IGameStateData = {
  game: GameData;
  loading: boolean;
  loadingMove: boolean;
  error: SerializedError;
  pickHint: boolean;
  page: EPage;
  loggedIn: boolean;
  leaderboard: ILeaderboard;
};
