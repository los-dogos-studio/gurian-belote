import type { Card } from '~/client/card'
import CardFace from '../CardFace'

interface TrumpSelectionPlayerPanelProps {
	cards: Card[]
	controls: React.ReactNode
}

const TrumpSelectionPlayerPanel = ({
	cards,
	controls,
}: TrumpSelectionPlayerPanelProps) => {
	return (
		<div className="flex items-center justify-center gap-8">
			<div className="flex justify-center items-center -space-x-4">
				{cards.map((card, index) => (
					<CardFace
						key={index}
						card={card}
						className="cursor-not-allowed pointer-events-none"
					/>
				))}
			</div>
			{controls}
		</div>
	)
}

export default TrumpSelectionPlayerPanel
