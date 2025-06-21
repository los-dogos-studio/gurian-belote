import { Card, Suit } from "../card";
import { PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import "reflect-metadata";
import { stringMapToIntEnumMap } from "./enum-map-utils";
import { Transform, Type } from "class-transformer";
import { IsEnum, ValidateNested } from "class-validator";
import { Trick } from "./trick";

export enum HandStage {
	TableTrumpSelection = "TableTrumpSelection",
	FreeTrumpSelection = "FreeTrumpSelection",
	HandInProgress = "HandInProgress",
	HandFinished = "HandFinished"
}

export abstract class HandState {
	@IsEnum(HandStage)
	state: HandStage;

	constructor(state: HandStage) {
		this.state = state;
	}
}

export class TableTrumpSelectionHandState extends HandState {
	@Type(() => Card)
	@ValidateNested()
	tableTrumpCard: Card

	@IsEnum(PlayerId)
	startingPlayer: PlayerId;

	@Type(() => Map<PlayerId, boolean>)
	@Transform(stringMapToIntEnumMap)
	selectionStatus: Map<PlayerId, boolean>


	constructor(
		state: HandStage,
		startingPlayer: PlayerId,
		tableTrumpCard: Card,
		selectionStatus: Map<PlayerId, boolean>
	) {
		super(state);
		this.startingPlayer = startingPlayer;
		this.tableTrumpCard = tableTrumpCard;
		this.selectionStatus = selectionStatus;
	}
}

export class FreeTrumpSelectionHandState extends HandState {
	@Type(() => Card)
	@ValidateNested()
	tableTrumpCard: Card

	@IsEnum(PlayerId)
	startingPlayer: PlayerId;

	@Type(() => Map<PlayerId, boolean>)
	@Transform(stringMapToIntEnumMap)
	selectionStatus: Map<PlayerId, boolean>

	constructor(
		state: HandStage,
		tableTrumpCard: Card,
		selectionStatus: Map<PlayerId, boolean>,
		startingPlayer: PlayerId
	) {
		super(state);
		this.startingPlayer = startingPlayer;
		this.tableTrumpCard = tableTrumpCard;
		this.selectionStatus = selectionStatus;
	}
}

export class InProgressHandState extends HandState {
	@IsEnum(Suit)
	trump: Suit

	@Type(() => Trick)
	@ValidateNested()
	trick: Trick

	@Type(() => Map<TeamId, number>)
	@Transform(stringMapToIntEnumMap)
	totals: Map<TeamId, number>

	constructor(
		state: HandStage,
		trump: Suit,
		trick: Trick,
		totals: Map<TeamId, number>
	) {
		super(state);
		this.trump = trump;
		this.trick = trick;
		this.totals = totals;
	}
}

export type HandStateType =
	| TableTrumpSelectionHandState
	| FreeTrumpSelectionHandState
	| InProgressHandState;

