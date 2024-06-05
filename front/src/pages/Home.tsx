import React from 'react';
import { Link } from 'react-router-dom';

const HomePage: React.FC = () => {
  return (
    <div className=" flex flex-col items-center justify-center bg-black text-white p-4">
      <h1 className="text-4xl font-bold text-center mb-6 tracking-wider">fivevfive</h1>
      <p className="text-lg text-center mb-6 max-w-md">
        fivevfive offers an easy and convenient way for you and your friends to create and manage custom games.
        Our system ensures that matches are balanced based on player ELO, and champion selection is influenced
        by mastery points, making every game competitive and fair.
      </p>
      <Link to="/games" className="bg-gray-700 hover:bg-gray-800 text-white font-bold py-2 px-4 rounded">
        Start a Game
      </Link>
      <p className="text-lg text-center mb-6 pt-6 max-w-md">
        Currently available only for me and my freinds, beacuse it would violate my riot api key eula if i made it public.
      </p>
    </div>
  );
};

export default HomePage;

