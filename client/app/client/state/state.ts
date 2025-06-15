import { IsArray, IsEnum, IsObject, IsOptional, IsString, ValidateNested } from "class-validator";
import type { Card } from "../card";
import { PlayerId } from "../player-id";
import { GameState } from "./game-state";
import { Type } from "class-transformer";
import "reflect-metadata";

export class State {
	@ValidateNested()
	@IsObject()
	@Type(() => GameState)
	gameState: GameState;

	@IsString()
	userId: string;

	@IsEnum(PlayerId)
	playerId: PlayerId;


	@IsOptional()
	@ValidateNested({ each: true })
	@IsArray()
	userCards?: Card[];

	constructor(
		gameState: GameState,
		userId: string,
		playerId: PlayerId,
		userCards: Card[]
	) {
		this.gameState = gameState;
		this.userId = userId;
		this.playerId = playerId;
		this.userCards = userCards;
	}
}
