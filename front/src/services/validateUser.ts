import { userbackendurl } from "@/common/constants";
import { getToken } from "./tokenUtils";

export default async function validateUser(username: string) {
  try {
    const body = { username: username };
    const response = await fetch(userbackendurl + "checkUser", {
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + getToken(),
      },
      method: "POST",
      body: JSON.stringify(body),
    });

    if (!response.ok) {
      return null;
    }

    return await response.json();
  } catch (error) {
    return null;
  }
}
