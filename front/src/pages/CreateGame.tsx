import { queryClient } from "@/utils/http";
import { PlayerSelector } from "@/components/PlayerSelector";
import { getAllUsers } from "@/services/getAllUsers";
import { useQuery } from "@tanstack/react-query";
export default function CreateGame() {

  const { data, isError, error } = useQuery({
    queryKey: ['users'],
    queryFn: getAllUsers,
    staleTime: 1000 * 60 * 60
  })

  if (!data) {
    return <div>Loading...</div>
  }


  if (isError) {
    return <div>Error: {error.message}</div>
  }


  return (
    <div className="">
      <PlayerSelector players={data} />
    </div>
  );
}

export const loader = () => {
  return queryClient.fetchQuery({
    queryKey: ['users'],
    queryFn: getAllUsers
  })
}
