import React from 'react';
import { Link } from 'react-router-dom';

const Header: React.FC = () => {
  return (
    <header className="w-full bg-black text-white py-4">
      <div className="max-w-screen-xl mx-auto px-4">
        <Link to="/" className='text-xl font-bold flex justify-center items-center hover:cursor-pointer' >
          fivevfive
        </Link>
      </div>
    </header>
  );
};

export default Header;
