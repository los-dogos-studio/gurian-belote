import { useGameState } from "../GameStateContext";

export const WaitingRoom = () => {
	const { gameState } = useGameState();

	return (
		<>
			<h1>
				Waiting Room
			</h1>
			{JSON.stringify(gameState)}
		</>
	);
}

export default WaitingRoom;
