import { User } from "../common/types";

export default async function getRemoteUserInfo(): Promise<User | null> {
  try {
    const response = await fetch("url", {
      headers: {
        "Content-Type": "application/json",
      },
      method: "GET",
      credentials: "include",
    });

    if (!response.ok) {
      return null;
    }

    const data: User = await response.json();
    return data;
  } catch (error) {
    return null;
  }
}
