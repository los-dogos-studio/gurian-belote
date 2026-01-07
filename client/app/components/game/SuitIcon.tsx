import { LuClub, LuDiamond, LuHeart, LuSpade } from 'react-icons/lu'
import { Suit } from '~/client/card'

interface SuitIconProps {
	suit: Suit
	iconClassName?: string
}

const SuitIcon = ({ suit, iconClassName = '' }: SuitIconProps) => {
	switch (suit) {
		case Suit.Spades:
			return <LuSpade className={iconClassName} />
		case Suit.Hearts:
			return <LuHeart className={iconClassName} />
		case Suit.Diamonds:
			return <LuDiamond className={iconClassName} />
		case Suit.Clubs:
			return <LuClub className={iconClassName} />
		default:
			throw new Error('Invalid suit')
	}
}

export default SuitIcon
