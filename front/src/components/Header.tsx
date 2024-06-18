import React from "react";
import { Link } from "react-router-dom";
import { VscAccount } from "react-icons/vsc";

const Header: React.FC = () => {
  return (
    <header className="w-full bg-black text-white py-4">
      <div className=" flex justify-between mx-auto px-4">
        <div></div>
        <Link to="/" className="text-xl font-bold hover:cursor-pointer">
          fivevfive
        </Link>
        <VscAccount className="text-2xl ml-4 hover:cursor-pointer" />
      </div>
    </header>
  );
};

export default Header;
