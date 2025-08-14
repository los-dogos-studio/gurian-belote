import { Card, Suit } from "../card";
import { getNextPlayerId, PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import "reflect-metadata";
import { stringMapToIntEnumMap } from "./enum-map-utils";
import { Transform, Type } from "class-transformer";
import { IsEnum, IsOptional, ValidateNested } from "class-validator";
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

	@Type(() => Trick)
	@IsOptional()
	@ValidateNested()
	previousTrick?: Trick;

	constructor(state: HandStage, previousTrick?: Trick) {
		this.state = state;
		this.previousTrick = previousTrick;
	}

	abstract getCurrentTurn(): PlayerId;

	abstract getPlayableCards(playerCards: Card[]): Card[];
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

	getCurrentTurn(): PlayerId {
		let result = this.startingPlayer;
		for (const [playerId, selected] of this.selectionStatus.entries()) {
			if (selected) {
				result = getNextPlayerId(playerId);
			}
		}
		return result;
	}

	getPlayableCards(_playerCards: Card[]): Card[] {
		return [];
	}
}

// TODO: DRY this code
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

	getCurrentTurn(): PlayerId {
		let result = this.startingPlayer;
		for (const [playerId, selected] of this.selectionStatus.entries()) {
			if (selected) {
				result = getNextPlayerId(playerId);
			}
		}
		return result;
	}

	getPlayableCards(_playerCards: Card[]): Card[] {
		return [];
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

	getCurrentTurn(): PlayerId {
		let result = this.trick.startingPlayer;
		for (let i = 0; i < this.trick.playedCards.size; i++) {
			result = getNextPlayerId(result);
		}
		return result;
	}

	getPlayableCards(playerCards: Card[]): Card[] {
		const trumpSuit = this.trump;
		const trick = this.trick;

		if (trick.playedCards.size === 0) return [...playerCards];

		const firstCard = trick.playedCards.get(trick.startingPlayer)!;
		const leadSuit = firstCard.suit;

		const cardsOfLeadSuit = playerCards.filter(card => card.suit === leadSuit);

		if (cardsOfLeadSuit.length > 0) return cardsOfLeadSuit;

		const playerTrumpCards = playerCards.filter(card => card.suit === trumpSuit);
		const trumpPlayed = Array.from(trick.playedCards.values()).some(card => card.suit === trumpSuit);

		if (trumpPlayed && playerTrumpCards.length > 0) {
			const highestTrump = Array.from(trick.playedCards.values())
				.filter(card => card.suit === trumpSuit)
				.reduce((highest, current) => current.compare(highest, trumpSuit) > 0 ? current : highest);

			const higherTrumpCards = playerTrumpCards.filter(card => card.compare(highestTrump, trumpSuit) > 0);
			return higherTrumpCards.length > 0 ? higherTrumpCards : playerTrumpCards;
		}

		return playerTrumpCards.length > 0 ? playerTrumpCards : [...playerCards];
	}
}

export type HandStateType =
	| TableTrumpSelectionHandState
	| FreeTrumpSelectionHandState
	| InProgressHandState;

