import { useGameState } from '../../GameContext'
import Scoreboard from '../../Scoreboard'
import PlayerPanel from '../../PlayerPanel'
import { getNextPlayerId, type PlayerId } from '~/client/player-id'
import { HandStage, InProgressHandState } from '~/client/state/hand'
import { TeamId } from '~/client/team-id'
import PlayerIcon from './PlayerIcon'
import TrumpDisplay from './TrumpDisplay'
import PlayArea from './PlayArea'

export const InGame = () => {
	const { gameState } = useGameState()

	if (!gameState) {
		return <div className="text-white">Waiting for game...</div>
	}

	const leftPlayerId: PlayerId = getNextPlayerId(gameState.playerId)
	const topPlayerId: PlayerId = getNextPlayerId(leftPlayerId)
	const rightPlayerId: PlayerId = getNextPlayerId(topPlayerId)

	const leftPlayerName =
		gameState.gameState.players.get(leftPlayerId) ??
		`Player ${leftPlayerId}`
	const topPlayerName =
		gameState.gameState.players.get(topPlayerId) ?? `Player ${topPlayerId}`
	const rightPlayerName =
		gameState.gameState.players.get(rightPlayerId) ??
		`Player ${rightPlayerId}`

	const currentPlayerId = gameState.gameState.hand?.getCurrentTurn()
	const currentPlayerOutline =
		'outline outline-1 outline-[#F3E5AB] shadow-[0_0_6px_rgba(243,229,171,0.6)]'

	const scores = {
		'Team 1': gameState.gameState.scores.get(TeamId.Team1) ?? -1,
		'Team 2': gameState.gameState.scores.get(TeamId.Team2) ?? -1,
	}

	return (
		<div className="h-full w-full relative gap-4 p-4 text-white">
			<div className="absolute top-2 right-2 p-3 flex gap-4">
				{gameState.gameState.hand &&
					gameState.gameState.hand.state ===
						HandStage.HandInProgress && (
						<TrumpDisplay
							suit={
								(
									gameState.gameState
										.hand as InProgressHandState
								).trump
							}
						/>
					)}
				<Scoreboard scores={scores} className="h-min" />
			</div>

			<div className="absolute top-1/2 left-2 transform -translate-y-1/2 p-3">
				<PlayerIcon
					label={leftPlayerName}
					className={
						currentPlayerId === leftPlayerId
							? currentPlayerOutline
							: ''
					}
				/>
			</div>

			<div className="absolute top-2 left-1/2 transform -translate-x-1/2 p-3">
				<PlayerIcon
					label={topPlayerName}
					className={
						currentPlayerId === topPlayerId
							? currentPlayerOutline
							: ''
					}
				/>
			</div>

			<div className="absolute top-1/2 right-2 transform -translate-y-1/2 p-3">
				<PlayerIcon
					label={rightPlayerName}
					className={
						currentPlayerId === rightPlayerId
							? currentPlayerOutline
							: ''
					}
				/>
			</div>

			<div className="absolute bottom-2 left-1/2 transform -translate-x-1/2 p-3">
				<PlayerPanel
					className={
						currentPlayerId === gameState.playerId
							? currentPlayerOutline
							: ''
					}
				/>
			</div>

			<div className="inline-block absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-60 h-80">
				<PlayArea />
			</div>
		</div>
	)
}

export default InGame
