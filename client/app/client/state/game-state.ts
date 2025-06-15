import { IsEnum, IsObject, IsOptional, IsString, ValidateNested } from "class-validator";
import type { PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import { Hand } from "./hand";
import { Type } from "class-transformer";
import 'reflect-metadata';

export enum GameStage {
	GameReady = "Ready",
	GameInProgress = "InProgress",
	GameFinished = "Finished"
}

export class GameState {
	@IsString()
	roomId: string;

	@Type(() => Map<PlayerId, string>)
	@ValidateNested()
	players: Map<PlayerId, string>;

	@Type(() => Map<TeamId, string[]>)
	@ValidateNested()
	teams: Map<TeamId, string[]>;

	@Type(() => Hand)
	@IsOptional()
	@ValidateNested()
	hand?: Hand;

	@IsEnum(GameStage)
	gameState: GameStage;

	@Type(() => Map<TeamId, number>)
	@ValidateNested()
	scores: Map<TeamId, number>;

	constructor(
		roomId: string,
		players: Map<PlayerId, string>,
		teams: Map<TeamId, string[]>,
		gameState: GameStage,
		scores: Map<TeamId, number>,
		hand?: Hand) {
		this.roomId = roomId;
		this.players = players;
		this.teams = teams;
		this.hand = hand;
		this.gameState = gameState;
		this.scores = scores;
	}
}
