import { gamebackendurl } from "@/common/constants";
import { Game } from "@/common/types";

interface GetGameParams {
  gameid: string;
  signal?: AbortSignal;
}

export const getChampions = async ({ gameid }: GetGameParams): Promise<Game> => {
  console.log(gameid)
  try {
    const resp = await fetch(gamebackendurl + 'games/' + gameid + "/champions", {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
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
