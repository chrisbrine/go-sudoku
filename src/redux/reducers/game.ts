import { PayloadAction } from "@reduxjs/toolkit";
import { IGameStateData } from "../types";
import { EPage } from "../../types";

export const setLoading = (
  state: IGameStateData,
  action: PayloadAction<boolean>
) => {
  state.loading = action.payload;
};

export const setPickHint = (
  state: IGameStateData,
  action: PayloadAction<boolean>
) => {
  state.pickHint = action.payload;
};

export const setPage = (
  state: IGameStateData,
  action: PayloadAction<EPage>
) => {
  state.page = action.payload;
};
