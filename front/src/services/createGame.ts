import { gamebackendurl } from "@/common/constants";

export const createGame = async (summoners: string[]): Promise<{ game_id: number }> => {
  try {

    const body = JSON.stringify({
      summoners: summoners,
      team_red: "gladiatorite",
      team_blue: "kaputite",
    })

    const resp = await fetch(gamebackendurl + 'games', {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: body
    })

    if (!resp.ok) {
      console.log('error fething all users')
    }

    const data = await resp.json()

    console.log(data)

    return data

  } catch (error) {
    console.log('error fething all users')
    return { game_id: 0 }
  }
}
