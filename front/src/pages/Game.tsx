import { Game as GameType, Summoner } from "@/common/types";
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
import { getNewChampion } from "@/services/getNewChampion";
import { swapSummoners } from "@/services/swapSummoners";

export default function Game() {
  const params = useParams();

  const [team1Swap, setTeam1Swap] = useState('')
  const [team2Swap, setTeam2Swap] = useState('')

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

  const { mutate: mutate3, } = useMutation({
    mutationFn: getNewChampion,
    onSuccess: (newData) => {
      queryClient.invalidateQueries({ queryKey: ['game', params.id] });
      setBlueTeam(newData.team_blue);
      setRedTeam(newData.team_red);
    }
  }
  );

  const { mutate: mutate4, } = useMutation({
    mutationFn: swapSummoners,
    onSuccess: (newData) => {
      queryClient.invalidateQueries({ queryKey: ['game', params.id] });
      setBlueTeam(newData.team_blue);
      setRedTeam(newData.team_red);
    }
  }
  );


  const roleOrder: { [key: string]: number } = {
    "top": 1,
    "jungle": 2,
    "mid": 3,
    "adc": 4,
    "support": 5
  };

  const sortSummonersByRole = (summoners: Summoner[]): Summoner[] => {
    return summoners.sort((a, b) => roleOrder[a.role] - roleOrder[b.role]);
  };

  const [blueTeam, setBlueTeam] = useState<GameType["team_blue"] | null>(null);
  const [redTeam, setRedTeam] = useState<GameType["team_red"] | null>(null);

  useEffect(() => {
    if (data) {
      setBlueTeam({
        ...data.team_blue,
        summoners: sortSummonersByRole(data.team_blue.summoners)
      });
      setRedTeam({
        ...data.team_red,
        summoners: sortSummonersByRole(data.team_red.summoners)
      });
    }
  }, [data]);


  useEffect(() => {
    if (team1Swap !== '' && team2Swap !== '') {
      mutate4({ team1Swap: team1Swap, team2Swap: team2Swap, gameid: params.id != undefined ? params.id : '0' })
      setTeam1Swap('')
      setTeam2Swap('')
    }
  }, [team1Swap, team2Swap]);

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

  const getNewChampFn = (user: string) => {
    mutate3({ username: user, gameid: params.id != undefined ? params.id : '0' })
  }


  return (
    <div className="flex flex-col gap-4 items-center">
      <div className="flex flex-col-3 ">
        <TeamTable team={blueTeam} newChamp={getNewChampFn} teamSwap={team1Swap} setTeamSwap={setTeam1Swap} />
        <RolesSplitter team={redTeam} />
        <TeamTable team={redTeam} newChamp={getNewChampFn} teamSwap={team2Swap} setTeamSwap={setTeam2Swap} />
      </div>
      <div className="flex flex-col-3 bg-black text-white">
        <Button className="w-32" onClick={getNewTeamsFn}>Reroll Teams</Button>
        <Button className="mx-4">Select Winner</Button>
        <Button className="w-32" onClick={getChamps}>Get Champions</Button>
      </div>
    </div>
  );
}




