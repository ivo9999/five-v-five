import React from "react";
import { Team } from "@/common/types";

interface TeamTableProps {
  team: Team;
  newChamp: (user: string) => void;
  teamSwap: string;
  setTeamSwap: (user: string) => void;
}

const TeamTable: React.FC<TeamTableProps> = ({
  team,
  newChamp,
  teamSwap,
  setTeamSwap,
}) => {
  return (
    <div className="bg-black text-white rounded-lg">
      <h2 className="text-2xl font-bold text-center mb-4">{team.name}</h2>
      <table className="table-auto w-full">
        <thead>
          <tr className="border-b-white border-b-2">
            {team.id % 2 === 0 && (
              <th className="px-4 w-44 py-2 text-center">Name</th>
            )}
            <th className="px-4 py-2 w-36 text-center">Champion</th>
            {team.id % 2 !== 0 && (
              <th className="px-4 py-2 w-44 text-center">Name</th>
            )}
          </tr>
        </thead>
        <tbody>
          {team.summoners.map((summoner) => (
            <tr key={summoner.id} className="border-b-white border-b-1">
              {team.id % 2 === 0 && (
                <td
                  onClick={() => setTeamSwap(summoner.name)}
                  className={`px-4 py-2 ${
                    teamSwap === summoner.name ? "bg-red-500 rounded" : ""
                  } text-center hover:cursor-pointer border-b-white border-b`}
                >
                  {decodeURIComponent(summoner.name)}
                </td>
              )}
              <td
                onClick={() => newChamp(summoner.name)}
                className="px-4 py-2 hover:cursor-pointer text-center border-b-white border-b"
              >
                {summoner.champion}
              </td>
              {team.id % 2 !== 0 && (
                <td
                  onClick={() => setTeamSwap(summoner.name)}
                  className={`px-4 py-2 ${
                    teamSwap === summoner.name ? "bg-red-500 rounded-xl" : ""
                  } text-center hover:cursor-pointer border-b-white border-b`}
                >
                  {decodeURIComponent(summoner.name)}
                </td>
              )}
            </tr>
          ))}
        </tbody>
      </table>
      <div className="flex flex-col items-center">
        <p>Team Mastery: {team.mastery_points}</p>
        <p>Team Rating: {team.rating}</p>
      </div>
    </div>
  );
};

export default TeamTable;
