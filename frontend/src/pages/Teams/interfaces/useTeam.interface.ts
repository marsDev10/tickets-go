import type { IApiResponseCreateTeam } from "./apiTeams.interface";

// Como podria representar el retorno de un array
export type IResponseHandleTeam = [
  error: string | boolean | null,
  data: IApiResponseCreateTeam | null
];
