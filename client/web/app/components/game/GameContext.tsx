import {
	createContext,
	type ReactNode,
	type RefObject,
	useContext,
	useEffect,
	useRef,
	useState,
} from 'react'
import { GameClient, type State } from '@gurian-belote/lib'
import { WebSessionManager } from '~/session/web-session-manager'

export interface GameStateContextType {
	gameState: State | null
	setGameState: (state: State | null) => void
}

const GameStateContext = createContext<GameStateContextType | null>(null)
const GameClientRefContext = createContext<RefObject<GameClient> | null>(null)

const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
const wsUrl = `${protocol}//${window.location.host}/ws`

export const GameProvider = ({ children }: { children: ReactNode }) => {
	const clientRef = useRef(new GameClient(wsUrl, new WebSessionManager()))
	const [gameState, setGameState] = useState<State | null>(null)

	useEffect(() => {
		const client = clientRef.current
		client.addListener(setGameState)
		client.restoreConnection()

		return () => {
			client.removeListener(setGameState)
		}
	}, [])

	return (
		<GameStateContext.Provider value={{ gameState, setGameState }}>
			<GameClientRefContext.Provider value={clientRef}>
				{children}
			</GameClientRefContext.Provider>
		</GameStateContext.Provider>
	)
}

export const useGameClient = () => {
	const context = useContext(GameClientRefContext)
	if (!context) {
		throw new Error('useGameClient must be used within a GameProvider')
	}
	return context.current
}

export const useGameState = () => {
	const context = useContext(GameStateContext)
	if (!context) {
		throw new Error('useGameState must be used within a GameProvider')
	}
	return context
}
