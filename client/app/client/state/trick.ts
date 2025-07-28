import { Transform, Type } from "class-transformer";
import type { Card } from "../card";
import { stringMapToIntEnumMap } from "./enum-map-utils";
import { IsEnum } from "class-validator";
import { PlayerId } from "../player-id";

export class Trick {
	@Type(() => Map<PlayerId, Card>)
	@Transform(stringMapToIntEnumMap)
	playedCards: Map<PlayerId, Card>

	@IsEnum(PlayerId)
	startingPlayer: PlayerId

	constructor(playedCards: Map<PlayerId, Card>, startingPlayer: PlayerId) {
		this.playedCards = playedCards;
		this.startingPlayer = startingPlayer;
	}
}

