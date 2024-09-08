import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { API } from "../api";
import { gamesReducer } from "./slices/game";

export const reducers = combineReducers({
  game: gamesReducer,
});

export const store = configureStore({
  reducer: reducers,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      thunk: {
        extraArgument: API,
      },
    }),
});
