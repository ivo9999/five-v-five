import { UserData } from "@/common/types";

export const getLocalUserData = (): UserData => {
  const storedUserString = localStorage.getItem("user");

  if (storedUserString) {
    const storedUser = JSON.parse(storedUserString);
    return storedUser;
  }
  return { username: "", id: 0 };
};
