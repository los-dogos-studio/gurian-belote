import { createContext, useContext, useEffect, useRef } from 'react';
import GameClient from '~/client/game-client';
import { useGameState } from './GameStateContext';

const GameClientContext = createContext<GameClient | null>(null);
const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const wsUrl = `${protocol}//${window.location.host}/ws`;

const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const wsUrl = `${protocol}//${window.location.host}/ws`;

export const GameClientProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const clientRef = useRef(new GameClient(wsUrl));
	const { setGameState } = useGameState();

	useEffect(() => {
		clientRef.current.addListener(setGameState);
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

