import Panel from "~/components/Panel";
import { useGameState } from "../GameStateContext";
import ListPanel from "~/components/ListPanel";
import { TeamId } from "~/client/team-id";
import Button from "~/components/Button";
import { useGameClient } from "../GameClientContext";
import { LuCopy, LuPlus, LuLogOut } from "react-icons/lu";

interface TeamColumnProps {
	teamId: TeamId,
	members: string[];
	onJoin: (teamId: TeamId) => void;
};

const TeamColumn = ({
	teamId,
	members,
	onJoin
}: TeamColumnProps) => {
	const JoinButton = () => (
		<Button
			onClick={() => onJoin(teamId)}
			variant="secondary"
		>
			<LuPlus className="inline-block mr-2" />
			Join
		</Button>
	);

	return (
		<ListPanel
			title={`Team ${teamId === TeamId.Team1 ? 1 : 2}`}
			items={members}
			footer={<JoinButton />}
			emptyLabel="No players yet."
		/>
	);
};

export const TeamSelection = () => {
	const { gameState } = useGameState();
	const client = useGameClient();

	if (!gameState || !gameState.gameState) {
		return <div>Invalid state...</div>; // TODO
	}

	const Title = () => {
		return (
			<h1 className="text-2xl font-bold text-amber-100/90 text-center mb-3 tracking-wider">
				Choose Your Team
			</h1>
		);
	}

	const RoomIdLabel = () => (
		<div className="mb-4 flex items-center justify-center space-x-2">
			<span className="text-sm font-bold text-amber-100/70 tracking-wider">
				{`Room ID: ${gameState.gameState.roomId}`}
			</span>
			<LuCopy
				className="inline-block cursor-pointer text-gray-400/30 hover:text-amber-100/70 transition-colors duration-200"
				onClick={() => navigator.clipboard.writeText(gameState.gameState.roomId)}
				size={"1.25rem"}
				title="Copy Room ID" />
		</div>
	);

	const handleJoinTeam = (teamId: TeamId) => {
		client.chooseTeam(teamId);
	}

	let joinedPlayersCount = 0;
	if (gameState.gameState.teams.has(TeamId.Team1)) {
		joinedPlayersCount += gameState.gameState.teams.get(TeamId.Team1)!.length;
	}
	if (gameState.gameState.teams.has(TeamId.Team2)) {
		joinedPlayersCount += gameState.gameState.teams.get(TeamId.Team2)!.length;
	}

	return (
		<div className="w-full h-full grid place-items-center">
			<Panel className="max-w-md w-full border border-yellow-600/30 rounded-2xl p-8">
				<Title />
				<RoomIdLabel />
				<div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full max-w-3xl mx-auto mb-6">
					<TeamColumn
						teamId={TeamId.Team1}
						members={gameState.gameState.teams.get(TeamId.Team1) ?? []}
						onJoin={handleJoinTeam}
					/>
					<TeamColumn
						teamId={TeamId.Team2}
						members={gameState.gameState.teams.get(TeamId.Team2) ?? []}
						onJoin={handleJoinTeam}
					/>
				</div>
				<Button
					onClick={() => client.startGame()}
					variant="primary"
					disabled={joinedPlayersCount !== 4}
					className="mb-4"
				>
					Start Game
				</Button>
				<Button
					onClick={() => client.leaveRoom()}
					variant="secondary"
				>
					<LuLogOut className="w-4 h-4 mr-2" />
					Leave Room
				</Button>
			</Panel>
		</div>
	);
}

export default TeamSelection;
