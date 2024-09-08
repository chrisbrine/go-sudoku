import { useDispatch, useSelector } from "react-redux";
import type { AppDispatch, RootState } from "./types";

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: <TSelected>(
  selector: (state: RootState) => TSelected,
  EqualityFn?: (left: TSelected, right: TSelected) => boolean
) => TSelected = useSelector;
