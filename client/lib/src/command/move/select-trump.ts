import type { Suit } from '../../card'

export class SelectTrumpMove {
	readonly command: string = 'selectTrump'
	suit: Suit | null

	constructor(suit: Suit | null) {
		this.suit = suit
	}
}

export default SelectTrumpMove
