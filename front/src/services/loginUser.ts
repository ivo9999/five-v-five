import { userbackendurl } from "@/common/constants";

export default async function loginUser(username: string, password: string) {
  const body = {
    username: username,
    password: password,
  };

  try {
    const res = await fetch(userbackendurl + "login", {
      headers: {
        "Content-Type": "application/json",
      },
      method: "POST",
      body: JSON.stringify(body),
      credentials: "include",
    });

    if (!res.ok) {
      return null;
    }

    return await res.json();
  } catch (error) {
    return null;
  }
}
