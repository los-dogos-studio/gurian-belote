import type { Card } from '../../card'

export class PlayCardMove {
	readonly command: string = 'playCard'
	card: Card
	skipDeclarations: boolean

	constructor(card: Card, skipDeclarations: boolean = false) {
		this.card = card
		this.skipDeclarations = skipDeclarations
	}
}

export default PlayCardMove
