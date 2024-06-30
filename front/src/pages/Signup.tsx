import React, { useEffect, useState } from "react";
import { Button } from "../components/ui/button";
import { Input as TextInput } from "../components/ui/input";
import { Link } from "react-router-dom";
import registerUser from "@/services/registerUser";

const SignupPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [leagueName, setLeagueName] = useState("");
  const [leagueTag, setLeagueTag] = useState("");
  const [discordName, setDiscordName] = useState("");
  const [password, setPassword] = useState("");

  const registerUserFunc = async () => {
    const userResp = await registerUser(
      username,
      password,
      leagueName,
      leagueTag,
      discordName,
    );
    console.log(userResp);
  };

  return (
    <div className="flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Sign Up</h1>
        <div className="space-y-4">
          <TextInput
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="w-full"
          />
          <TextInput
            placeholder="League Name"
            value={leagueName}
            onChange={(e) => setLeagueName(e.target.value)}
            className="w-full"
          />
          <TextInput
            placeholder="League Tag"
            value={leagueTag}
            onChange={(e) => setLeagueTag(e.target.value)}
            className="w-full"
          />
          <TextInput
            placeholder="Discord Name"
            value={discordName}
            onChange={(e) => setDiscordName(e.target.value)}
            className="w-full"
          />
          <TextInput
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            type="password"
            className="w-full"
          />
          <Button className="w-full" onClick={registerUserFunc}>
            Sign Up
          </Button>
        </div>
        <p className="mt-6 text-center text-light">
          Already have an account?{" "}
          <Link to="/login" className="text-white underline">
            Log In
          </Link>
        </p>
      </div>
    </div>
  );
};

export default SignupPage;
