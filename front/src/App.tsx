import { BrowserRouter, Routes, Route, Outlet } from "react-router-dom";
import { QueryClientProvider } from "@tanstack/react-query";
import LoginPage from "./pages/Login"; // Adjust path as necessary
import SignupPage from "./pages/Signup"; // Adjust path as necessary
import HomePage from "./pages/Home";
import CreateGame, { loader as CreateGameLoader } from "./pages/CreateGame";
import Header from "./components/Header";
import Game from "./pages/Game";
import { queryClient } from "./utils/http";
import { loader as GameLoader } from "./utils/loaders/loadGame";
import NotFoundPage from "./pages/NotFound";
import { AuthProvider } from "./contexts/AuthContext";

function App() {
  const Root = () => {
    return (
      <div className="flex flex-col bg-black h-screen">
        <Header />
        <div className="flex-1 flex items-center justify-center">
          <Outlet />
        </div>
      </div>
    );
  };

  return (
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Root />}>
              <Route index element={<HomePage />} />
              <Route path="*" element={<NotFoundPage />} />
              <Route path="login" element={<LoginPage />} />
              <Route path="signup" element={<SignupPage />} />
              <Route
                path="games"
                element={<CreateGame />}
                loader={CreateGameLoader}
              />
              <Route path="games/:id" element={<Game />} loader={GameLoader} />
            </Route>
          </Routes>
        </BrowserRouter>
      </QueryClientProvider>
    </AuthProvider>
  );
}

export default App;
