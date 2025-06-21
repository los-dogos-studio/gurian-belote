export enum PlayerId {
	NoId = 0,
	Player1 = 1,
	Player2 = 2,
	Player3 = 3,
	Player4 = 4
}

export function getNextPlayerId(playerId: PlayerId): PlayerId {
	return (playerId - 1 + 1) % 4 + 1;
}

export function getPreviousPlayerId(playerId: PlayerId): PlayerId {
	return (playerId - 1 + 3) % 4 + 1;
}
