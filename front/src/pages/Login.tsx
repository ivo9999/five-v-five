import React from 'react';
import { Button, } from '../components/ui/button';
import { Input as TextInput } from '../components/ui/input'; const LoginPage: React.FC = () => {

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Login</h1>
        <div className="space-y-4">
          <TextInput placeholder="Email" className="w-full" />
          <TextInput placeholder="Password" type="password" className="w-full" />
          <Button className="w-full">Login</Button>
        </div>
        <p className="mt-6 text-center text-light">
          Don’t have an account? <a href="/signup" className="text-white underline">Sign up</a>
        </p>
      </div>
    </div>
  );
};

export default LoginPage;
