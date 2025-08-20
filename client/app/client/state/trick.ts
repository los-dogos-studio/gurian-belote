import { Transform, Type } from "class-transformer";
import { Card } from "../card";
import { enumKeyMapToClassValue } from "./enum-map-utils";
import { IsEnum } from "class-validator";
import { PlayerId } from "../player-id";

export class Trick {
	@Type(() => Map<PlayerId, Card>)
	@Transform(enumKeyMapToClassValue(Card))
	playedCards: Map<PlayerId, Card>

	@IsEnum(PlayerId)
	startingPlayer: PlayerId

	constructor(playedCards: Map<PlayerId, Card>, startingPlayer: PlayerId) {
		this.playedCards = playedCards;
		this.startingPlayer = startingPlayer;
	}
}

