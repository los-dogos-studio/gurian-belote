import { useState } from 'react'
import type { Card } from '@gurian-belote/lib'
import CardFace from '../CardFace'
import { useGameClient } from '../GameContext'
import SkipDeclarationsToggle from './SkipDeclarationsToggle'

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
	const [skipDeclarations, setSkipDeclarations] = useState(false)

	const onCardClick = (card: Card) => {
		gameClient.playCard(card, skipDeclarations)
	}

	return (
		<div className="flex flex-col items-center gap-3 w-full">
			<SkipDeclarationsToggle
				active={skipDeclarations}
				disabled={disabled}
				onToggle={() => setSkipDeclarations((prev) => !prev)}
			/>

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
		</div>
	)
}

export default PlayerCardsPanel
