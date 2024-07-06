import React from "react";
import { Link } from "react-router-dom";
import { VscAccount } from "react-icons/vsc";
import { useAuth } from "@/hooks/useAuth";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

const Header: React.FC = () => {
  const { currentUser, logOutUser } = useAuth();

  return (
    <header className="w-full bg-black text-white py-4">
      <div className=" flex justify-between mx-auto px-4">
        <div style={{ visibility: currentUser ? "visible" : "hidden" }}>
          {currentUser && (
            <button className="opacity-0" aria-hidden="true">
              <VscAccount className="text-2xl hover:cursor-default ml-4" />
            </button>
          )}
          {!currentUser && <button className="opacity-0">Sign Up</button>}
        </div>
        <Link to="/" className="text-xl font-bold hover:cursor-pointer">
          fivevfive
        </Link>
        {!currentUser && <Link to="/signup">Sign Up</Link>}
        {currentUser && (
          <DropdownMenu>
            <DropdownMenuTrigger>
              <VscAccount className="text-2xl ml-4 hover:cursor-pointer" />
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>Profile</DropdownMenuItem>
              <DropdownMenuItem onClick={() => logOutUser()}>
                Log out
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        )}
      </div>
    </header>
  );
};

export default Header;
