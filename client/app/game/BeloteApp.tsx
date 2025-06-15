import { GameClientProvider } from "./GameClientContext";
import { GameStateProvider } from "./GameStateContext";
import GamePageDispatcher from "./GamePageDispatcher";

export const BeloteApp = () => {
	return (
		<GameStateProvider>
			<GameClientProvider>
				<GamePageDispatcher />
			</GameClientProvider>
		</GameStateProvider>
	);
}

export default BeloteApp;
