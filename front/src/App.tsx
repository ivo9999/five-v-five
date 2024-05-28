import { BrowserRouter, Routes, Route, Outlet } from 'react-router-dom';
import LoginPage from './pages/Login'; // Adjust path as necessary
import SignupPage from './pages/Signup'; // Adjust path as necessary
import HomePage from './pages/Home';
import CreateGame from './pages/CreateGame';
import Header from './components/Header';

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
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Root />}>
          <Route index element={<HomePage />} />
          <Route path="login" element={<LoginPage />} />
          <Route path="signup" element={<SignupPage />} />
          <Route path="games" element={<CreateGame />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
