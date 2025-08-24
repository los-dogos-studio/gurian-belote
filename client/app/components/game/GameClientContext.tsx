import { createContext, useContext, useEffect, useRef } from 'react';
import GameClient from '~/client/game-client';
import { SessionManager } from '~/client/session-manager';
import { useGameState } from './GameStateContext';

const GameClientContext = createContext<GameClient | null>(null);
const wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const httpProtocol = window.location.protocol === 'https:' ? 'https:' : 'http:';
const wsUrl = `${wsProtocol}//${window.location.host}/ws`;
const authUrl = `${httpProtocol}//${window.location.host}/auth`;

export const GameClientProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
	const clientRef = useRef(new GameClient(wsUrl, authUrl, new SessionManager()));
	const { setGameState } = useGameState();

	useEffect(() => {
		clientRef.current.addListener(setGameState);
		
		const sessionInfo = clientRef.current.getStoredSession();
		
		if (sessionInfo?.roomId && sessionInfo?.autoReconnect) {
			clientRef.current.reconnect({
				onError: (error) => console.warn('Auto-reconnection failed:', error)
			});
		}
		
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

