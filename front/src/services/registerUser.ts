export default async function registerUser(
  username: string,
  password: string,
  leagueName: string,
  leagueTag: boolean,
  discordName: string,
) {
  const body = {
    username: username,
    league_name: leagueName,
    league_tag: leagueTag,
    password: password,
    discord_name: discordName,
  };

  try {
    const res = await fetch("url", {
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
      body: JSON.stringify(body),
    });

    if (!res.ok) {
      return null;
    }

    return await res.json();
  } catch (error) {
    console.log("Error:", error);
    return null;
  }
}
