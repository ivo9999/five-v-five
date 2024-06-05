import { getGame } from "@/services/getGame";
import { queryClient } from "../http";
import { Game } from "@/common/types";
import { LoaderFunction } from "react-router-dom";

type LoaderParams = {
  id: string;
}

export const loader: LoaderFunction<LoaderParams> = ({ params }) => {
  return queryClient.fetchQuery<Game | null>({
    queryKey: ['game', params.id],
    queryFn: ({ signal }) => getGame({ signal, gameid: params.id != undefined ? params.id : '0' }),
  });
};


