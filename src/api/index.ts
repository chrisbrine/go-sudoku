import {
  GameData,
  ILeaderboard,
  ILogin,
  IRegister,
  IRowColNum,
  IUpdateDifficulty,
  IUpdateName,
  IUpdatePassword,
  IUpdateUserName,
  LoginData,
} from "../types";

export class API {
  private static port = 3000;
  private static host = window.location.hostname;
  private static token = "";
  private static username = "";

  private static handleToken = () => {
    API.token = localStorage.getItem("token") || "";
    API.username = localStorage.getItem("username") || "";
  };
  private static getToken = () => {
    if (!API.token) {
      this.handleToken();
    }
    return API.token;
  };
  private static getUsername = () => {
    if (!API.username) {
      this.handleToken();
    }
    return API.username;
  };
  private static setHandle = () => {
    if (API.token) {
      localStorage.setItem("token", API.token);
    }
    if (API.username) {
      localStorage.setItem("username", API.username);
    }
  };
  private static setToken = (token: string) => {
    API.token = token;
    this.setHandle();
  };
  private static setUsername = (username: string) => {
    API.username = username;
    this.setHandle();
  };
  private static clearToken = () => {
    API.token = "";
    API.username = "";
    localStorage.removeItem("token");
    localStorage.removeItem("username");
  };

  private static url = () => `http://${API.host}:${API.port}`;

  public static async handle<T = unknown, T2 = unknown>(
    method: string,
    path: string,
    data?: T2
  ): Promise<T> {
    const reqURL = `${API.url()}${path}`;
    const token = this.getToken();
    const username = this.getUsername();
    const request = {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: token ? `Bearer ${token}` : "",
        "X-Username": username,
      },
      body: data ? JSON.stringify(data) : undefined,
    };
    const response = await fetch(reqURL, request);

    if (response.ok) {
      return (await response.json()) as T;
    } else {
      throw new Error(response.statusText);
    }
  }

  public static hasToken(): boolean {
    return !!this.getToken();
  }

  public static async get(): Promise<GameData> {
    return API.handle<GameData>("GET", "/api/game");
  }

  public static async new(): Promise<GameData> {
    return API.handle<GameData>("GET", "/api/game/new");
  }

  public static async quit(): Promise<GameData> {
    return API.handle<GameData>("GET", "/api/game/quit");
  }

  public static async move({ col, row, num }: IRowColNum): Promise<GameData> {
    return API.handle<GameData>("POST", `/api/game/move/${row}/${col}/${num}`);
  }

  public static async hint({ col, row, num }: IRowColNum): Promise<GameData> {
    return API.handle<GameData>("POST", `/api/game/hint/${row}/${col}/${num}`);
  }

  public static async hintRemove({
    col,
    row,
    num,
  }: IRowColNum): Promise<GameData> {
    return API.handle<GameData>(
      "POST",
      `/api/game/hintRemove/${row}/${col}/${num}`
    );
  }

  public static async login(
    username: string,
    password: string
  ): Promise<GameData> {
    const response = await API.handle<LoginData, ILogin>("POST", "/api/login", {
      username,
      password,
    });
    this.setToken(response.token);
    this.setUsername(username);

    return API.get();
  }

  public static async logout(): Promise<void> {
    await API.handle<void>("POST", "/api/logout");
    this.clearToken();
  }

  public static async register(
    username: string,
    password: string,
    name: string
  ): Promise<GameData> {
    const response = await API.handle<LoginData, IRegister>(
      "POST",
      "/api/register",
      {
        username,
        password,
        name,
      }
    );

    this.setToken(response.token);
    this.setUsername(username);
    return API.get();
  }

  public static async updateUserName(username: string): Promise<GameData> {
    await API.handle<void, IUpdateUserName>("POST", "/api/update/username", {
      username,
    });

    API.username = username;
    return API.get();
  }

  public static async updatePassword(
    oldPassword: string,
    newPassword: string
  ): Promise<GameData> {
    await API.handle<void, IUpdatePassword>("POST", "/api/update/password", {
      oldPassword,
      newPassword,
    });

    return API.get();
  }

  public static async updateName(name: string): Promise<GameData> {
    await API.handle<void, IUpdateName>("POST", "/api/update/name", {
      name,
    });

    return API.get();
  }

  public static async updateDifficulty(difficulty: number): Promise<GameData> {
    await API.handle<void, IUpdateDifficulty>(
      "GET",
      `/api/update/difficulty/${difficulty}`
    );

    return API.get();
  }

  public static async leaderboard(): Promise<ILeaderboard> {
    return API.handle<ILeaderboard>("GET", "/api/game/leaderboard");
  }
}
