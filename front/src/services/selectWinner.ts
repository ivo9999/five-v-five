import { gamebackendurl } from "@/common/constants";

export const selectWinner = async (gameid: number, winner: string) => {
  try {
    const resp = await fetch(gamebackendurl + "games/" + gameid + "/winner", {
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
      body: JSON.stringify({ winner: winner }),
    });

    if (!resp.ok) {
      return null;
    }

    return await resp.json();
  } catch (error) {
    return null;
  }
};
