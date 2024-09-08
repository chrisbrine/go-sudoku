import type { RootState } from "../types";

export const getGame = (state: RootState) => state.game.game;
export const getBoard = (state: RootState) => state.game.game.Board;
export const getHints = (state: RootState) => state.game.game.Hints;
export const getNumbersLeft = (state: RootState) => state.game.game.NumbersLeft;
export const getPlaying = (state: RootState) => state.game.game.Playing;
export const getGameStatus = (state: RootState) => state.game.game.GameStatus;
export const getLastMove = (state: RootState) => state.game.game.LastMove;
export const getMistakes = (state: RootState) => state.game.game.Mistakes;
export const getWins = (state: RootState) => state.game.game.Wins;
export const getLosses = (state: RootState) => state.game.game.Losses;
export const getPerfectWins = (state: RootState) => state.game.game.PerfectWins;
export const getPoints = (state: RootState) => state.game.game.Points;
export const getDifficulty = (state: RootState) => {
  const difficulty = state.game.game.Difficulty;
  if (difficulty < 1 || difficulty > 3) {
    return 1;
  }
  return difficulty;
};
export const getPickHint = (state: RootState) => state.game.pickHint;
export const getLoading = (state: RootState) => state.game.loading;
export const getLoadingMove = (state: RootState) => state.game.loadingMove;
export const getError = (state: RootState) => state.game.error;
export const getPage = (state: RootState) => state.game.page;
export const getLoggedIn = (state: RootState) => state.game.loggedIn;
export const getInGame = (state: RootState) => state.game.game.InGame;

export const getUserName = (state: RootState) => state.game.game.Username;
export const getName = (state: RootState) => state.game.game.Name;

export const getLeaderboard = (state: RootState) => state.game.leaderboard;
