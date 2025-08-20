import { Card, Suit } from "../card";
import { getNextPlayerId, PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import "reflect-metadata";
import { enumKeyMap } from "./enum-map-utils";
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
	@Transform(enumKeyMap)
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
	@Transform(enumKeyMap)
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
	@Transform(enumKeyMap)
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
		return playerCards.filter(card => this.isCardPlayable(card, playerCards));
	}

	private isCardPlayable(card: Card, playerCards: Card[]): boolean {
		if (!playerCards.includes(card)) {
			return false;
		}

		if (this.trick.playedCards.size === 0) {
			return true;
		}

		const firstCard = this.trick.playedCards.get(this.trick.startingPlayer)!;
		const leadSuit = firstCard.suit;

		if (this.hasCardOfSuit(leadSuit, playerCards)) {
			var requiredSuit = leadSuit;
		} else if (this.hasCardOfSuit(this.trump, playerCards)) {
			var requiredSuit = this.trump;
		} else {
			return true;
		}

		if (card.suit !== requiredSuit) {
			return false;
		}

		if (requiredSuit === this.trump) {
			const highestTrumpInTrick = this.getHighestTrumpInTrick();
			const highestPlayerTrump = playerCards
				.filter(c => c.suit === this.trump)
				.sort((a, b) => -a.compare(b, this.trump))[0];

			if (highestTrumpInTrick && highestPlayerTrump && highestPlayerTrump.compare(highestTrumpInTrick, this.trump) > 0) {
				return card.compare(highestTrumpInTrick, this.trump) > 0;
			}
		}

		return true;
	}

	private hasCardOfSuit(suit: Suit, playerCards: Card[]): boolean {
		return playerCards.some(card => card.suit === suit);
	}

	private getHighestTrumpInTrick(): Card | null {
		let highestTrump: Card | null = null;
		for (const card of this.trick.playedCards.values()) {
			if (card.suit === this.trump) {
				if (!highestTrump || card.compare(highestTrump, this.trump) > 0) {
					highestTrump = card;
				}
			}
		}
		return highestTrump;
	}
}

export type HandStateType =
	| TableTrumpSelectionHandState
	| FreeTrumpSelectionHandState
	| InProgressHandState;

