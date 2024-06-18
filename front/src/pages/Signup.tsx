import React from "react";
import { Button } from "../components/ui/button";
import { Input as TextInput } from "../components/ui/input";
import { Link } from "react-router-dom";

const SignupPage: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Sign Up</h1>
        <div className="space-y-4">
          <TextInput placeholder="Username" className="w-full" />
          <TextInput placeholder="Email" className="w-full" />
          <TextInput placeholder="League Name" className="w-full" />
          <TextInput placeholder="League Tag" className="w-full" />
          <TextInput
            placeholder="Discord Name"
            type="password"
            className="w-full"
          />
          <TextInput
            placeholder="Password"
            type="password"
            className="w-full"
          />
          <TextInput
            placeholder="Confirm Password"
            type="password"
            className="w-full"
          />
          <Button className="w-full">Sign Up</Button>
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
