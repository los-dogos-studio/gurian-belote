import { IsEnum, IsOptional, IsString, ValidateNested } from "class-validator";
import type { PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import { Transform, Type } from "class-transformer";
import 'reflect-metadata';
import { Hand } from "./hand";
import { stringMapToIntEnumMap } from "./enum-map-utils";

export enum GameStage {
	GameReady = "Ready",
	GameInProgress = "InProgress",
	GameFinished = "Finished"
}

export class GameState {
	@IsString()
	roomId: string;

	@Type(() => Map<PlayerId, string>)
	@Transform(stringMapToIntEnumMap)
	players: Map<PlayerId, string>;

	@Type(() => Map<TeamId, string[]>)
	@Transform(stringMapToIntEnumMap)
	teams: Map<TeamId, string[]>;

	@Type(() => Hand)
	@IsOptional()
	@ValidateNested()
	hand?: Hand;

	@IsEnum(GameStage)
	gameState: GameStage;

	@Type(() => Map<TeamId, number>)
	@Transform(stringMapToIntEnumMap)
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
