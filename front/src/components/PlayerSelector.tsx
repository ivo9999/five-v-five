import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

const players: Player[] = [
  {
    summonerName: "Neko",
    division: "Iron",
  },
  {
    summonerName: "yass",
    division: "Bronze",
  },
  {
    summonerName: "zed",
    division: "Gold",
  }
]

type Player = {
  summonerName: string
  division: string
}

export default function PlayerSelector() {
  return (
    <Table className="bg-black text-white">
      <TableCaption>Players</TableCaption>
      <TableHeader >
        <TableRow className="hover:bg-black">
          <TableHead className="w-[200px]">Summoner Name</TableHead>
          <TableHead>Division</TableHead>
          <TableHead>Remove</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {players.map((player) => (
          <TableRow key={player.summonerName} className="hover:bg-zinc-800">
            <TableCell className="font-medium text-xl">{player.summonerName}</TableCell>
            <TableCell className="font-medium text-xl">{player.division}</TableCell>
            <TableCell className="text-center text-xl hover:cursor-pointer">X</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  )
}

