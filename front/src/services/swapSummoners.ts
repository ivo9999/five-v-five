import { gamebackendurl } from "@/common/constants";
import { Game } from "@/common/types";

interface GetGameParams {
  gameid: string;
  team1Swap: string;
  team2Swap: string;
  signal?: AbortSignal;
}

export const swapSummoners = async ({ gameid, team1Swap, team2Swap }: GetGameParams): Promise<Game> => {
  try {

    const body = JSON.stringify({ summoner1: team1Swap, summoner2: team2Swap })

    const resp = await fetch(gamebackendurl + 'games/' + gameid + "/swap", {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: body
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
