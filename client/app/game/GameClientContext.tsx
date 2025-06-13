import { createContext, useContext, useEffect, useRef } from 'react';
import GameClient from '~/client/game-client';

const GameClientContext = createContext<GameClient | null>(null);

export const GameClientProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const clientRef = useRef(new GameClient("ws://localhost:8080/ws")); // TODO: Make this configurable

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

