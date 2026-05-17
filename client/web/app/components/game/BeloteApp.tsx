import { GameProvider } from './GameContext'
import GamePageDispatcher from './GamePageDispatcher'
import Background from '../Background'

export const BeloteApp = () => {
	return (
		<GameProvider>
			<Background>
				<GamePageDispatcher />
			</Background>
		</GameProvider>
	)
}

export default BeloteApp
