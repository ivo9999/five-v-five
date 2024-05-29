export type Player = {
  summonerName: string
  division: string
}

export type User = {
  username: string,
  league_name: string,
  league_tag: string,
  discord_name: string,
  Id: number
}

export interface Summoner {
  name: string;
  role: string;
  champion: string;
  id: number;
  team_id: number;
}

export interface Team {
  name: string;
  summoners: Summoner[];
  id: number;
  rating: number;
  mastery_points: number;
}

export interface Game {
  winner: string;
  date: string;
  team_blue: Team;
  team_red: Team;
  id: number;
}
