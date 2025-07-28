import { createContext, useContext, useEffect, useRef } from 'react';
import GameClient from '~/client/game-client';
import { useGameState } from './GameStateContext';

const GameClientContext = createContext<GameClient | null>(null);

export const GameClientProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const clientRef = useRef(new GameClient("ws://localhost:8080/ws")); // TODO: Make this configurable
	const { setGameState } = useGameState();

	clientRef.current.addListener(setGameState);

	useEffect(() => {
		return () => {
			clientRef.current.disconnect();
		};
	}, []);

	return (
		<GameClientContext.Provider value={clientRef.current}>
			{children}
		</GameClientContext.Provider>
	);
};

export const useGameClient = () => {
	const ctx = useContext(GameClientContext);
	if (!ctx) {
		throw new Error('useGameClient must be used within a GameClientProvider');
	}
	return ctx;
};

