export const updateUserData = (username: string, id: number) => {
  const updatedUser = { username, id };
  localStorage.setItem("user", JSON.stringify(updatedUser));
};
