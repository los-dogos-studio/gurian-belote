import { TeamId } from './team-id'

export enum PlayerId {
	NoId = 0,
	Player1 = 1,
	Player2 = 2,
	Player3 = 3,
	Player4 = 4,
}

export function getNextPlayerId(playerId: PlayerId): PlayerId {
	return ((playerId - 1 + 1) % 4) + 1
}

export function getPreviousPlayerId(playerId: PlayerId): PlayerId {
	return ((playerId - 1 + 3) % 4) + 1
}

export function getTeamId(playerId: PlayerId): TeamId {
	if (playerId === PlayerId.Player1 || playerId === PlayerId.Player3)
		return TeamId.Team1
	if (playerId === PlayerId.Player2 || playerId === PlayerId.Player4)
		return TeamId.Team2
	return TeamId.NoTeam
}
