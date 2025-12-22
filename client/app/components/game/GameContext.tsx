import { createContext, useContext, useEffect, useRef, useState } from 'react';
import GameClient from '~/client/game-client';
import type { State } from "~/client/state/state";

export interface GameStateContextType {
	gameState: State | null;
	setGameState: (state: State | null) => void;
}

const GameStateContext = createContext<GameStateContextType | null>(null);
const GameClientContext = createContext<GameClient | null>(null);

const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const wsUrl = `${protocol}//${window.location.host}/ws`;

export const GameProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const clientRef = useRef(new GameClient(wsUrl));
	const [gameState, setGameState] = useState<State | null>(null);

	useEffect(() => {
		clientRef.current.addListener(setGameState);
		return () => {
			clientRef.current.disconnect();
		};
	}, []);

	return (
		<GameStateContext.Provider value={{ gameState, setGameState }}>
			<GameClientContext.Provider value={clientRef.current}>
				{children}
			</GameClientContext.Provider>
		</GameStateContext.Provider>
	);
};

export const useGameClient = () => {
	const context = useContext(GameClientContext);
	if (!context) {
		throw new Error('useGameClient must be used within a GameProvider');
	}
	return context;
}

export const useGameState = () => {
	const context = useContext(GameStateContext);
	if (!context) {
		throw new Error('useGameState must be used within a GameProvider');
	}
	return context;
}
