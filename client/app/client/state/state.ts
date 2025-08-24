import { IsArray, IsEnum, IsObject, IsOptional, IsString, ValidateNested } from "class-validator";
import { Card } from "../card";
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
	token: string;

	@IsEnum(PlayerId)
	playerId: PlayerId;

	@Type(() => Card)
	@ValidateNested({ each: true })
	@IsOptional()
	@IsArray()
	userCards?: Card[];

	constructor(
		gameState: GameState,
		token: string,
		playerId: PlayerId,
		userCards: Card[]
	) {
		this.gameState = gameState;
		this.token = token;
		this.playerId = playerId;
		this.userCards = userCards;
	}
}
