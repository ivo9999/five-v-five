import { gamebackendurl } from "@/common/constants";
import { Game } from "@/common/types";

interface GetGameParams {
  gameid: string;
  signal?: AbortSignal;
}

export const getGame = async ({ gameid, signal }: GetGameParams): Promise<Game> => {
  try {
    const resp = await fetch(gamebackendurl + 'games/' + gameid, {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
      signal
    })


    if (!resp.ok) {
      console.log('error fething game')
    }

    const data = await resp.json()

    return data

  } catch (error) {
    console.log('error fething game')
    throw error
  }
}
