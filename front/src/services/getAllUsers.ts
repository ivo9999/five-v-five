import { userbackendurl } from "@/common/constants";
import { User } from "@/common/types";

export const getAllUsers = async (): Promise<User[]> => {
  try {
    const resp = await fetch(userbackendurl + 'getAllUsers', {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'GET',
    })

    if (!resp.ok) {
      console.log('error fething all users')
    }

    const data = await resp.json()

    console.log(data)

    return data

  } catch (error) {
    console.log('error fething all users')
    return []
  }
}
