import React, { useState } from "react";
import { Button } from "../components/ui/button";
import { Input as TextInput } from "../components/ui/input";
import { Link } from "react-router-dom";
import loginUser from "@/services/loginUser";
const LoginPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const loginUserFunc = async () => {
    const loginRes = await loginUser(username, password);
    console.log(loginRes);
  };

  return (
    <div className="flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Log In</h1>
        <div className="space-y-4">
          <TextInput
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="w-full"
          />
          <TextInput
            placeholder="Password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full"
          />
          <Button className="w-full" onClick={loginUserFunc}>
            Log In
          </Button>
        </div>
        <p className="mt-6 text-center text-light">
          Don’t have an account?{" "}
          <Link to="/signup" className="text-white underline">
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
};

export default LoginPage;
