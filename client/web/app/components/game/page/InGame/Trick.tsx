import type { Card } from '@gurian-belote/lib'
import CardFace from '../../CardFace'

export enum CardPosition {
	Left = 'left',
	Top = 'top',
	Right = 'right',
	Bottom = 'bottom',
}

interface TrickProps {
	bottom?: Card
	left?: Card
	top?: Card
	right?: Card
	startingPosition?: CardPosition
}

const CardSlot = ({
	card,
	className = '',
}: {
	card?: Card
	className?: string
}) => {
	return <div className={className}>{card && <CardFace card={card} />}</div>
}

const clockwiseCardSlotOrder = [
	CardPosition.Left,
	CardPosition.Top,
	CardPosition.Right,
	CardPosition.Bottom,
]
const cardSlotZClassNames = ['z-10', 'z-20', 'z-30', 'z-40']

function getZIndexClasses(lowestZCard: CardPosition = CardPosition.Left) {
	const lowestZIndex = clockwiseCardSlotOrder.indexOf(lowestZCard)
	const zClasses = {} as Record<CardPosition, string>

	clockwiseCardSlotOrder.forEach((pos, index) => {
		const zIndex =
			(index - lowestZIndex + clockwiseCardSlotOrder.length) %
			clockwiseCardSlotOrder.length
		zClasses[pos] = cardSlotZClassNames[zIndex]
	})

	return zClasses
}

const Trick = ({ bottom, left, top, right, startingPosition }: TrickProps) => {
	const zClasses = getZIndexClasses(startingPosition)

	return (
		<div className="relative w-full h-full">
			<CardSlot
				card={bottom}
				className={`absolute bottom-2 left-1/2 transform -translate-x-1/2 ${zClasses[CardPosition.Bottom]}`}
			/>
			<CardSlot
				card={left}
				className={`absolute left-2 top-1/2 transform -translate-y-1/2 ${zClasses[CardPosition.Left]}`}
			/>
			<CardSlot
				card={top}
				className={`absolute top-2 left-1/2 transform -translate-x-1/2 ${zClasses[CardPosition.Top]}`}
			/>
			<CardSlot
				card={right}
				className={`absolute right-2 top-1/2 transform -translate-y-1/2 ${zClasses[CardPosition.Right]}`}
			/>
		</div>
	)
}

export default Trick
