import { GameClientProvider } from "./GameClientContext";
import { GameStateProvider } from "./GameStateContext";
import GamePageDispatcher from "./GamePageDispatcher";
import Background from "../Background";

export const BeloteApp = () => {
	return (
		<GameStateProvider>
			<GameClientProvider>
				<Background>
					<GamePageDispatcher />
				</Background>
			</GameClientProvider>
		</GameStateProvider>
	);
}

export default BeloteApp;
