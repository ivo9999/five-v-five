import PlayerSelector from "../components/PlayerSelector";
export default function CreateGame() {
  return (
    <div className="">
      <h1 className="text-4xl text-white">Game</h1>
      <div className="bg-background text-foreground">
        <PlayerSelector />
      </div>
    </div>
  );
}

