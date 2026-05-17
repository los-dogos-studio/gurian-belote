import { Suit } from '@gurian-belote/lib'
import Panel from '../../Panel'
import { useGameState } from '../GameContext'
import { GameStage } from '@gurian-belote/lib'
import {
	HandStage,
	InProgressHandState,
	TableTrumpSelectionHandState,
} from '@gurian-belote/lib'
import TableTrumpSelectionPlayerPanel from './TableTrumpSelectionPlayerPanel'
import FreeTrumpSelectionPlayerPanel from './FreeTrumpSelectionPlayerPanel'
import PlayerCardsPanel from './PlayerCardsPanel'

const PlayerPanelContent = () => {
	const { gameState } = useGameState()
	if (
		!gameState ||
		!gameState.gameState ||
		!gameState.gameState.hand ||
		!gameState.gameState.gameState ||
		gameState.gameState.gameState !== GameStage.GameInProgress
	) {
		return <div>Invalid game stage...</div>
	}

	const hand = gameState.gameState.hand
	const handState = hand.state
	const currentPlayerId = gameState.gameState.hand?.getCurrentTurn()
	const controlsDisabled = currentPlayerId !== gameState.playerId

	const skippable = true // TODO

	let trump = null
	if (
		handState === HandStage.TableTrumpSelection ||
		handState === HandStage.FreeTrumpSelection
	) {
		trump = (hand as TableTrumpSelectionHandState).tableTrumpCard.suit
	} else {
		trump = (hand as InProgressHandState).trump
	}
	const cards = gameState.userCards!.sort((a, b) => {
		if (a.suit !== b.suit) {
			const suitOrder = [
				Suit.Spades,
				Suit.Hearts,
				Suit.Diamonds,
				Suit.Clubs,
			]
			return suitOrder.indexOf(a.suit) - suitOrder.indexOf(b.suit)
		}
		return b.compare(a, trump)
	})

	switch (handState) {
		case HandStage.TableTrumpSelection:
			return (
				<TableTrumpSelectionPlayerPanel
					disabled={controlsDisabled}
					cards={cards}
				/>
			)
		case HandStage.FreeTrumpSelection:
			return (
				<FreeTrumpSelectionPlayerPanel
					disabled={controlsDisabled}
					forbiddenSuit={trump}
					cards={cards}
					skippable={skippable}
				/>
			)
		case HandStage.HandInProgress:
			return (
				<PlayerCardsPanel
					disabled={controlsDisabled}
					cards={cards}
					playableCards={hand.getPlayableCards(cards!)}
				/>
			)
		default:
			return <div>Invalid hand state...</div>
	}
}

export const PlayerPanel = ({ className = '' }: { className?: string }) => {
	return (
		<Panel className={`flex justify-center p-4 mb-2 gap-4 ${className}`}>
			<PlayerPanelContent />
		</Panel>
	)
}

export default PlayerPanel
