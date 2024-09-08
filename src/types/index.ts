export enum EPage {
  LOGIN = "login",
  REGISTER = "register",
  GAME = "game",
  DASHBOARD = "dashboard",
  PROFILE = "profile",
  LEADERBOARD = "leaderboard",
}

export enum EGameStatus {
  LOSE = -1,
  IN_PROGRESS = 0,
  WON = 1,
}

export enum ELastMoveType {
  NONE = "UNKNOWN",
  MOVE = "MOVE",
  HINT = "HINT",
  HINT_REMOVE = "HINT_REMOVE",
}

export interface ILastMove {
  Row: number;
  Col: number;
  Num: number;
  Type: ELastMoveType;
}

export interface GameData {
  Board: number[][];
  Hints: boolean[][][];
  NumbersLeft: number[];
  Mistakes: number;
  InGame: boolean;
  Playing: boolean;
  GameStatus: EGameStatus;
  LastMove: ILastMove;
  Username: string;
  Name: string;
  Wins: number;
  Losses: number;
  Points: number;
  PerfectWins: number;
  Difficulty: number;
  Success: boolean;
}

export interface LoginData {
  token: string;
}

export interface ILogin {
  username: string;
  password: string;
}

export interface IRegister {
  username: string;
  password: string;
  name: string;
}

export interface IUpdateUserName {
  username: string;
}

export interface IUpdatePassword {
  oldPassword: string;
  newPassword: string;
}

export interface IUpdateName {
  name: string;
}

export interface IUpdateDifficulty {
  difficulty: number;
}

export interface IRowCol {
  row: number;
  col: number;
}

export interface IRowColNum extends IRowCol {
  num: number;
}

export interface ILeader {
  Username: string;
  Name: string;
  Wins: number;
  Losses: number;
  Points: number;
  PerfectWins: number;
}

export type ILeaderboard = ILeader[];
