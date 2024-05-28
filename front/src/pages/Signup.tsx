import React from 'react';
import { Button, } from '../components/ui/button';
import { Input as TextInput } from '../components/ui/input';

const SignupPage: React.FC = () => {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-black text-white">
      <div className="w-full max-w-md">
        <h1 className="mb-6 text-3xl font-bold text-center">Sign Up</h1>
        <div className="space-y-4">
          <TextInput placeholder="Email" className="w-full" />
          <TextInput placeholder="Password" type="password" className="w-full" />
          <TextInput placeholder="Confirm Password" type="password" className="w-full" />
          <Button className="w-full">Sign Up</Button>
        </div>
        <p className="mt-6 text-center text-light">
          Already have an account? <a href="/login" className="text-white underline">Log In</a>
        </p>
      </div>
    </div>
  );
};

export default SignupPage;
