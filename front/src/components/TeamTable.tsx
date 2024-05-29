import { Team } from "@/common/types";

interface TeamTableProps {
  team: Team;
}
const TeamTable: React.FC<TeamTableProps> = ({ team }) => {
  return (
    <div className="bg-black text-white rounded-lg">
      <h2 className="text-2xl font-bold text-center mb-4">{team.name}</h2>
      <table className="table-auto w-full">
        <thead>
          <tr className="border-b-white border-b-2">
            {team.id % 2 == 0 && <th className="px-4 py-2 text-center">Name</th>}
            <th className="px-4 py-2 text-left">Champion</th>
            {team.id % 2 !== 0 && <th className="px-4 py-2 text-center">Name</th>}
          </tr>
        </thead>
        <tbody>
          {team.summoners.map(summoner => (
            <tr key={summoner.id} className="border-b-white border-b-1">
              {team.id % 2 == 0 && <td className="px-4 py-2  text-center border-b-white border-b">{decodeURIComponent(summoner.name)}</td>}
              <td className="px-4 py-2 border-b-white border-b">{summoner.champion}</td>
              {team.id % 2 !== 0 && <td className="px-4 py-2  text-center border-b-white border-b">{decodeURIComponent(summoner.name)}</td>}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TeamTable;
