import type { Card } from '~/client/card'
import CardFace from '../../CardFace'

interface TrickProps {
	bottom: Card | undefined
	left: Card | undefined
	top: Card | undefined
	right: Card | undefined
}

const CardSlot = ({
	card,
	className = '',
}: {
	card: Card | undefined
	className?: string
}) => {
	return <div className={className}>{card && <CardFace card={card} />}</div>
}

const Trick = ({ bottom, left, top, right }: TrickProps) => {
	// rethink?
	// TODO: add z
	return (
		<div className="relative w-full h-full">
			<CardSlot
				card={bottom}
				className="absolute bottom-2 left-1/2 transform -translate-x-1/2"
			/>
			<CardSlot
				card={left}
				className="absolute left-2 top-1/2 transform -translate-y-1/2"
			/>
			<CardSlot
				card={top}
				className="absolute top-2 left-1/2 transform -translate-x-1/2"
			/>
			<CardSlot
				card={right}
				className="absolute right-2 top-1/2 transform -translate-y-1/2"
			/>
		</div>
	)
}

export default Trick
