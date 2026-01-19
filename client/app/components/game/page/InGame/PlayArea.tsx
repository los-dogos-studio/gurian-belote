import {
	FreeTrumpSelectionHandState,
	HandStage,
	TableTrumpSelectionHandState,
	type InProgressHandState,
} from '~/client/state/hand'
import { useGameState } from '../../GameContext'
import Trick, { CardPosition } from './Trick'
import TableTrumpQuery from './TableTrumpQuery'
import { getNextPlayerId, PlayerId } from '~/client/player-id'

const PlayArea = () => {
	const { gameState } = useGameState()

	if (!gameState) {
		return <div className="text-white">Waiting for game...</div>
	}

	if (!gameState.gameState.hand) {
		return <div />
	}

	switch (gameState.gameState.hand.state) {
		case HandStage.TableTrumpSelection:
			return TableTrumpQuery(
				(gameState.gameState.hand as TableTrumpSelectionHandState)
					.tableTrumpCard,
				'Accept Table Trump?'
			)
		case HandStage.FreeTrumpSelection:
			return TableTrumpQuery(
				(gameState.gameState.hand as FreeTrumpSelectionHandState)
					.tableTrumpCard,
				'Select Table Trump'
			)

		case HandStage.HandInProgress: {
			const inProgressHand = gameState.gameState
				.hand as InProgressHandState
			const trick = inProgressHand.trick

			const leftPlayerId: PlayerId = getNextPlayerId(gameState.playerId)
			const topPlayerId: PlayerId = getNextPlayerId(leftPlayerId)
			const rightPlayerId: PlayerId = getNextPlayerId(topPlayerId)

			let startingPosition = undefined
			switch (trick.startingPlayer) {
				case leftPlayerId:
					startingPosition = CardPosition.Left
					break
				case topPlayerId:
					startingPosition = CardPosition.Top
					break
				case rightPlayerId:
					startingPosition = CardPosition.Right
					break
				case gameState.playerId:
					startingPosition = CardPosition.Bottom
					break
				default:
					startingPosition = undefined
			}

			return (
				<Trick
					bottom={trick.playedCards.get(gameState.playerId)}
					left={trick.playedCards.get(leftPlayerId)}
					top={trick.playedCards.get(topPlayerId)}
					right={trick.playedCards.get(rightPlayerId)}
					startingPosition={startingPosition}
				/>
			)
		}

		default:
			return <div />
	}
}

export default PlayArea
