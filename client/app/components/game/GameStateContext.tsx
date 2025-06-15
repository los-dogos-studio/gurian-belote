import { createContext, useContext, useState } from "react";
import type { State } from "~/client/state/state";

export interface GameStateContextType {
	gameState: State | null;
	setGameState: (state: State | null) => void;
}

const GameStateContext = createContext<GameStateContextType | null>(null);

export const GameStateProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const [gameState, setGameState] = useState<State | null>(null);

	return (
		<GameStateContext.Provider value={{gameState, setGameState}}>
			{children}
		</GameStateContext.Provider>
	);
}

export const useGameState = () => {
	const context = useContext(GameStateContext);
	if (!context) {
		throw new Error('useGameState must be used within a GameStateProvider');
	}
	return context;

}
