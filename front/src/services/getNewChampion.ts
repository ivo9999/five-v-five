import { gamebackendurl } from "@/common/constants";
import { Game } from "@/common/types";

interface GetGameParams {
  gameid: string;
  username: string;
  signal?: AbortSignal;
}

export const getNewChampion = async ({ gameid, username }: GetGameParams): Promise<Game> => {
  try {

    console.log(gameid, username)

    const body = JSON.stringify({ summoner_name: username })

    const resp = await fetch(gamebackendurl + 'games/' + gameid + "/newChamp", {
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
