import { GameStage } from "~/client/state/game-state";
import { useGameState } from "./GameStateContext";
import GameOver from "./page/GameOver";
import InGame from "./page/InGame";
import Lobby from "./page/Lobby";
import WaitingRoom from "./page/WaitingRoom";

export const GamePageDispatcher = () => {
	const { gameState } = useGameState();

	if (!gameState || !gameState.gameState) {
		return <Lobby />;
	}

	switch (gameState.gameState.gameState) {
		case GameStage.GameReady:
			return <WaitingRoom />
		case GameStage.GameInProgress:
			return <InGame />
		case GameStage.GameFinished:
			return <GameOver />
		default:
			return <div>Unknown game stage</div>;
	}
}

export default GamePageDispatcher;
