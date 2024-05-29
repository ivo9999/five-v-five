import { Game as GameType } from "@/common/types";
import { getGame } from "@/services/getGame"
import { queryClient } from "@/utils/http"
import { useMutation, useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import TeamTable from "@/components/TeamTable";
import RolesSplitter from "@/components/RolesSplitter";
import { Button } from "@/components/ui/button";
import { getChampions } from "@/services/getChampions";
import { getNewTeams } from "@/services/getNewTeams";

export default function Game() {
  const params = useParams();

  const { data, isError, error } = useQuery<GameType>({
    queryKey: ['game', params.id],
    queryFn: ({ signal }) => getGame({ signal, gameid: params.id != undefined ? params.id : '0' }),
  });

  const { mutate, } = useMutation({
    mutationFn: () => getChampions({ gameid: params.id != undefined ? params.id : '0' }),
    onSuccess: (newData) => {
      queryClient.invalidateQueries({ queryKey: ['game', params.id] });
      setBlueTeam(newData.team_blue);
      setRedTeam(newData.team_red);
    }
  }
  );

  const { mutate: mutate2, } = useMutation({
    mutationFn: () => getNewTeams({ gameid: params.id != undefined ? params.id : '0' }),
    onSuccess: (newData) => {
      queryClient.invalidateQueries({ queryKey: ['game', params.id] });
      setBlueTeam(newData.team_blue);
      setRedTeam(newData.team_red);
    }
  }
  );

  const [blueTeam, setBlueTeam] = useState<GameType["team_blue"] | null>(null);
  const [redTeam, setRedTeam] = useState<GameType["team_red"] | null>(null);

  useEffect(() => {
    if (data) {
      setBlueTeam(data.team_blue);
      setRedTeam(data.team_red);
    }
  }, [data]);

  if (isError) {
    return <div>Error: {error?.message}</div>;
  }

  if (!data || !blueTeam || !redTeam) {
    return <div>No data available</div>;
  }

  const getChamps = () => {
    mutate()
  }

  const getNewTeamsFn = () => {
    mutate2()
  }

  console.log(data)

  return (
    <div className="flex flex-col gap-4 items-center">
      <div className="flex flex-row ">
        <TeamTable team={blueTeam} />
        <RolesSplitter team={redTeam} />
        <TeamTable team={redTeam} />
      </div>
      <div className="flex flex-row bg-black text-white">
        <Button onClick={getNewTeamsFn}>Reroll Teams</Button>
        <Button className="bg-white text-black mx-6"></Button>
        <Button onClick={getChamps}>Get Champions</Button>
      </div>
    </div>
  );
}



export const loader = ({ params }: { params: { id: string } }) => {
  return queryClient.fetchQuery<GameType | null>({
    queryKey: ['game', params.id],
    queryFn: ({ signal }) => getGame({ signal, gameid: params.id }),
  });
};


