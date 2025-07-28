import { IsEnum, IsOptional, IsString, ValidateNested } from "class-validator";
import type { PlayerId } from "../player-id";
import type { TeamId } from "../team-id";
import { Transform, Type } from "class-transformer";
import 'reflect-metadata';
import { stringMapToIntEnumMap } from "./enum-map-utils";
import { FreeTrumpSelectionHandState, HandStage, HandState, InProgressHandState, TableTrumpSelectionHandState, type HandStateType } from "./hand";

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

	@Type(() => HandState, {
		keepDiscriminatorProperty: true,
		discriminator: {
			property: 'state',
			subTypes: [
				{ value: InProgressHandState, name: HandStage.HandInProgress },
				{ value: TableTrumpSelectionHandState, name: HandStage.TableTrumpSelection },
				{ value: FreeTrumpSelectionHandState, name: HandStage.FreeTrumpSelection }
			]
		}
	})
	@IsOptional()
	@ValidateNested()
	hand?: HandStateType;

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
		hand?: HandStateType) {
		this.roomId = roomId;
		this.players = players;
		this.teams = teams;
		this.hand = hand;
		this.gameState = gameState;
		this.scores = scores;
	}
}
