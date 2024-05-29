import React from 'react';
import { Button, } from '../components/ui/button';
import { Input as TextInput } from '../components/ui/input'; import { Link } from 'react-router-dom';
const LoginPage: React.FC = () => {

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Log In</h1>
        <div className="space-y-4">
          <TextInput placeholder="Email" className="w-full" />
          <TextInput placeholder="Password" type="password" className="w-full" />
          <Button className="w-full">Log In</Button>
        </div>
        <p className="mt-6 text-center text-light">
          Don’t have an account? <Link to="/signup" className="text-white underline">Sign up</Link>
        </p>
      </div>
    </div>
  );
};

export default LoginPage;
