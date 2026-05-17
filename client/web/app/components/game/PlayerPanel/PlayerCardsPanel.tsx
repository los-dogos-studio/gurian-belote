import type { Card } from '@gurian-belote/lib'
import CardFace from '../CardFace'
import { useGameClient } from '../GameContext'

interface PlayerCardsPanelProps {
	cards: Card[]
	playableCards: Card[]
	disabled: boolean
}

const PlayerCardsPanel = ({
	cards,
	playableCards,
	disabled,
}: PlayerCardsPanelProps) => {
	const gameClient = useGameClient()

	const onCardClick = (card: Card) => {
		gameClient.playCard(card)
	}

	return (
		<div className="flex justify-center items-center -space-x-4">
			{cards.map((card, index) => {
				const notPlayable = !playableCards.some((playableCard) =>
					playableCard.equals(card)
				)
				const cardClassName =
					disabled || notPlayable
						? 'brightness-70 cursor-not-allowed pointer-events-none'
						: ''
				return (
					<CardFace
						key={index}
						card={card}
						onClick={
							disabled || notPlayable ? () => {} : onCardClick
						}
						hover
						className={cardClassName}
					/>
				)
			})}
		</div>
	)
}

export default PlayerCardsPanel
