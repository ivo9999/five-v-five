import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Input } from "./ui/input";
import React, { useEffect, useState } from "react";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "./ui/button";
import { useNavigate } from "react-router-dom";
import { User } from "@/common/types";
import { createGame } from "@/services/createGame";

interface PlayerSelectorProps {
  players: User[];
}

export const PlayerSelector: React.FC<PlayerSelectorProps> = ({ players }) => {
  const [searchSummoner, setSearchSummoner] = useState("");
  const [selectedPlayers, setSelectedPlayers] = useState<User[]>([]);
  const [availablePlayers, setAvailablePlayers] = useState(players);
  const [filteredSummoners, setFilteredSummoners] = useState(availablePlayers);
  const [startError, setStartError] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    setFilteredSummoners(
      availablePlayers.filter((player) =>
        player.league_name.toLowerCase().includes(searchSummoner.toLowerCase()),
      ),
    );
  }, [searchSummoner]);

  const selectPlayer = (player: User) => {
    setSelectedPlayers([...selectedPlayers, player]);
    setAvailablePlayers(
      availablePlayers.filter((p) => p.league_name !== player.league_name),
    );
    setFilteredSummoners(
      availablePlayers.filter((p) => p.league_name !== player.league_name),
    );
  };

  const deselectPlayer = (player: User) => {
    setSelectedPlayers(
      selectedPlayers.filter((p) => p.league_name !== player.league_name),
    );
    setAvailablePlayers([...availablePlayers, player]);
    setFilteredSummoners([...availablePlayers, player]);
  };

  const StartGame = async () => {
    if (selectedPlayers.length < 2 || selectedPlayers.length > 10) {
      setStartError(true);
      return;
    }
    const resp = await createGame(
      selectedPlayers.map((player) => player.league_name),
    );
    navigate("/games/" + resp.game_id);
  };

  return (
    <>
      <div
        className="flex flex-col  justify-between"
        style={{ height: "calc(100vh - 60px)" }}
      >
        <Table className="bg-black mt-10 text-white">
          <TableHeader>
            <TableRow className="hover:bg-black">
              <TableHead className="w-[200px] text-lg">Summoner Name</TableHead>
              <TableHead className="text-lg">Username</TableHead>
              <TableHead className="text-lg">Remove</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="min-h-52">
            {selectedPlayers.map((player) => (
              <TableRow key={player.league_name} className="hover:bg-black">
                <TableCell className="font-medium hover:cursor-default text-xl">
                  {replaceAll(player.league_name, "%20", " ")}
                </TableCell>
                <TableCell className="font-medium hover:cursor-default text-xl">
                  {player.username}
                </TableCell>
                <TableCell>
                  <div
                    className="w-5 ml-6 text-xl text-center hover:cursor-pointer"
                    onClick={() => deselectPlayer(player)}
                  >
                    X
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <div className="relative grid grid-cols-3 grid-rows-2 justify-items-center w-full">
          <div
            className={`text-white text-xl ${startError ? "text-red-500" : ""}`}
          >
            {selectedPlayers.length}
          </div>
          <Popover>
            <PopoverTrigger asChild>
              <Button
                className="bg-black pb-3 text-white hover:text-white hover:bg-zinc-900"
                variant="outline"
              >
                +
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-100 bg-black border-black">
              <ScrollArea className="w-96 h-96 bg-black text-white rounded-md border">
                <div className="p-4">
                  <Input
                    type="text"
                    className="bg-black text-lg h-10 text-white mb-6"
                    placeholder="Search Summoner"
                    value={searchSummoner}
                    onChange={(e) => setSearchSummoner(e.target.value)}
                  />
                  {filteredSummoners.length > 0 &&
                    filteredSummoners.map((player) => (
                      <div key={player.league_name}>
                        <div
                          className="text-md hover:cursor-pointer"
                          onClick={() => selectPlayer(player)}
                        >
                          {replaceAll(player.league_name, "%20", " ")} |{" "}
                          {player.username}
                        </div>
                        <Separator className="my-2 text-white bg-white" />
                      </div>
                    ))}
                  {filteredSummoners.length === 0 && (
                    <div className="text-md text-center">
                      No summoners found
                    </div>
                  )}
                </div>
              </ScrollArea>
            </PopoverContent>
          </Popover>
          <Button
            className="ml-4 bg-black text-white hover:bg-zinc-900 hover:text-white border-white"
            variant="outline"
          >
            Import Discord
          </Button>
          <div></div>

          <Button
            onClick={StartGame}
            className="bg-white mb-14 text-black hover:bg-zinc-200 text-2xl"
          >
            Start
          </Button>
        </div>
      </div>
    </>
  );
};

function replaceAll(str: string, search: string, replacement: string): string {
  const escapedSearch = search.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  const regExp = new RegExp(escapedSearch, "g");
  return str.replace(regExp, replacement);
}
