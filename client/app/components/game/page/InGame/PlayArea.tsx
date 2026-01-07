import {
	FreeTrumpSelectionHandState,
	HandStage,
	TableTrumpSelectionHandState,
	type InProgressHandState,
} from '~/client/state/hand'
import { useGameState } from '../../GameContext'
import Trick from './Trick'
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

	const leftPlayerId: PlayerId = getNextPlayerId(gameState.playerId)
	const topPlayerId: PlayerId = getNextPlayerId(leftPlayerId)
	const rightPlayerId: PlayerId = getNextPlayerId(topPlayerId)

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
			return (
				<Trick
					bottom={inProgressHand.trick.playedCards.get(
						gameState.playerId
					)}
					left={inProgressHand.trick.playedCards.get(leftPlayerId)}
					top={inProgressHand.trick.playedCards.get(topPlayerId)}
					right={inProgressHand.trick.playedCards.get(rightPlayerId)}
				/>
			)
		}

		default:
			return <div />
	}
}

export default PlayArea
