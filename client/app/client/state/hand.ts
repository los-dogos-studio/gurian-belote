import { IsEnum, IsOptional, ValidateNested } from "class-validator";
import { Suit, type Card } from "../card";
import { PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import type { Trick } from "./trick";
import { Type } from "class-transformer";
import "reflect-metadata";

export enum HandState {
	TableTrumpSelection = "TableTrumpSelection",
	FreeTrumpSelection = "FreeTrumpSelection",
	HandInProgress = "HandInProgress",
	HandFinished = "HandFinished"
}

export class Hand {
	@IsEnum(HandState)
	state: HandState;

	@ValidateNested()
	@IsOptional()
	currentTrick?: Trick;

	@IsEnum(PlayerId)
	@IsOptional()
	startingPlayer?: PlayerId;

	@ValidateNested()
	@Type(() => Map<TeamId, Number>)
	totals: Map<TeamId, number>;

	@ValidateNested()
	@Type(() => Map<Card, Boolean>)
	playerCards: Map<PlayerId, Map<Card, boolean>>;

	@ValidateNested()
	tableTrumpCard: Card;

	@ValidateNested()
	@Type(() => Map<PlayerId, Boolean>)
	tableTrumpSelectionStatus: Map<PlayerId, boolean>;

	@ValidateNested()
	@Type(() => Map<PlayerId, Boolean>)
	freeTrumpSelectionStatus: Map<PlayerId, boolean>;

	@IsEnum(Suit)
	trump: Suit;

	constructor(
		state: HandState,
		currentTrick: Trick | undefined,
		startingPlayer: PlayerId | undefined,
		totals: Map<TeamId, number>,
		playerCards: Map<PlayerId, Map<Card, boolean>>,
		tableTrumpCard: Card,
		tableTrumpSelectionStatus: Map<PlayerId, boolean>,
		freeTrumpSelectionStatus: Map<PlayerId, boolean>,
		trump: Suit
	) {
		this.state = state;
		this.currentTrick = currentTrick;
		this.startingPlayer = startingPlayer;
		this.totals = totals;
		this.playerCards = playerCards;
		this.tableTrumpCard = tableTrumpCard;
		this.tableTrumpSelectionStatus = tableTrumpSelectionStatus;
		this.freeTrumpSelectionStatus = freeTrumpSelectionStatus;
		this.trump = trump;
	}
}
