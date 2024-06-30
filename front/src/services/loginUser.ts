import { userbackendurl } from "@/common/constants";
import { User } from "../common/types";

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

    const data: User = await res.json();
    return data;
  } catch (error) {
    console.log("Error:", error);
    return null;
  }
}
