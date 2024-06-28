import { useAuth } from "@/hooks/useAuth";
import { Navigate, useLocation } from "react-router-dom";

export type AuthenticatedProps = {
  children: React.ReactNode;
};

export default function Authenticated({
  children,
}: AuthenticatedProps): JSX.Element {
  const location = useLocation();
  const { currentUser, isLoading } = useAuth();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!currentUser) {
    return <Navigate replace to="/login" state={{ from: location }} />;
  }

  return <>{children}</>;
}
