import { Team } from "@/common/types";

interface TeamTableProps {
  team: Team;
}
const RolesSplitter: React.FC<TeamTableProps> = ({ team }) => {
  return (
    <div className="bg-black text-white  rounded-lg shadow-md">
      <h2 className="text-2xl font-bold text-black mb-4">text</h2>
      <table className="table-auto w-full">
        <thead>
          <tr className="border-b-white border-b-2">
            <th className="px-4 py-2 text-center">Role</th>
          </tr>
        </thead>
        <tbody>
          {team.summoners.map(summoner => (
            <tr key={summoner.id} className=" transition-colors">
              <td className="px-4 py-2 text-center border-b-white border-b">{summoner.role}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RolesSplitter;
