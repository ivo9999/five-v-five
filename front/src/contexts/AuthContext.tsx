import { createContext, useState, useEffect } from "react";
import { AuthContextType, AuthProviderProps, User } from "../common/types";
import { getLocalUserData } from "@/services/setLocalUserData";
import { updateLocalUserData } from "@/services/updateLocalUserData";
import getRemoteUserInfo from "@/services/getRemoteUserData";

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchUser = async () => {
      const user = getLocalUserData();
      if (user && user.username !== "") {
        try {
          const userInfo = await getRemoteUserInfo();
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

  return (
    <AuthContext.Provider
      value={{ currentUser, isLoading, logOutUser, setCurrentUser }}
    >
      {children}
    </AuthContext.Provider>
  );
};
