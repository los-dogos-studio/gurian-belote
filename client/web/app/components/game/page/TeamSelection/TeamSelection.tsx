import Panel from '~/components/Panel'
import { useGameClient, useGameState } from '../../GameContext'
import { TeamId } from '@gurian-belote/lib'
import Button from '~/components/Button'
import TeamColumn from './TeamColumn'
import RoomIdLabel from './RoomIdLabel'

const Title = () => {
	return (
		<h1 className="text-2xl font-bold text-amber-100/90 text-center mb-3 tracking-wider">
			Choose Your Team
		</h1>
	)
}

export const TeamSelection = () => {
	const { gameState } = useGameState()
	const client = useGameClient()

	if (!gameState || !gameState.gameState) {
		return <div>Invalid state...</div>
	}

	const handleJoinTeam = (teamId: TeamId) => {
		client.chooseTeam(teamId)
	}

	let joinedPlayersCount = 0
	if (gameState.gameState.teams.has(TeamId.Team1)) {
		joinedPlayersCount += gameState.gameState.teams.get(
			TeamId.Team1
		)!.length
	}
	if (gameState.gameState.teams.has(TeamId.Team2)) {
		joinedPlayersCount += gameState.gameState.teams.get(
			TeamId.Team2
		)!.length
	}

	return (
		<div className="w-full h-full grid place-items-center">
			<Panel className="max-w-md w-full border border-yellow-600/30 rounded-2xl p-8">
				<Title />
				<RoomIdLabel roomId={gameState.gameState.roomId} />
				<div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full max-w-3xl mx-auto mb-6">
					<TeamColumn
						teamId={TeamId.Team1}
						members={
							gameState.gameState.teams.get(TeamId.Team1) ?? []
						}
						onJoin={handleJoinTeam}
					/>
					<TeamColumn
						teamId={TeamId.Team2}
						members={
							gameState.gameState.teams.get(TeamId.Team2) ?? []
						}
						onJoin={handleJoinTeam}
					/>
				</div>
				<Button
					onClick={() => client.startGame()}
					variant="primary"
					disabled={joinedPlayersCount !== 4}
				>
					Start Game
				</Button>
			</Panel>
		</div>
	)
}

export default TeamSelection
