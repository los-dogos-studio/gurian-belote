import type { Card, Suit } from "../card";
import type { PlayerId } from "../player-id";

export interface Trick {
	startingPlayer: PlayerId;
	cards: Map<PlayerId, Card>;
	trump: Suit;
}
