import { createContext, useState, useEffect } from "react";
import { AuthContextType, AuthProviderProps, UserData } from "../common/types";
import { getLocalUserData } from "@/services/getLocalUserData";
import { updateLocalUserData } from "@/services/updateLocalUserData";
import validateUser from "@/services/validateUser";

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [currentUser, setCurrentUser] = useState<UserData | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchUser = async () => {
      const user = getLocalUserData();
      if (user && user.username !== "") {
        try {
          const userInfo = await validateUser(user.username);
          setCurrentUser(userInfo);
        } catch (error) {
          console.error("Failed to fetch user info:", error);
          setCurrentUser(null);
        }
      } else {
        logOutUser();
      }
      setIsLoading(false);
    };

    fetchUser();
  }, []);

  const logOutUser = () => {
    setCurrentUser(null);
    updateLocalUserData("", 0);
  };

  const logInUser = (userData: UserData) => {
    setCurrentUser(userData);
    updateLocalUserData(userData.username, userData.id);
  };

  return (
    <AuthContext.Provider
      value={{ currentUser, isLoading, logOutUser, logInUser }}
    >
      {children}
    </AuthContext.Provider>
  );
};
