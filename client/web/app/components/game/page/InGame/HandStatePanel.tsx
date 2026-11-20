import Panel from '~/components/Panel'
import {
	TeamId,
	PlayerId,
	DeclarationType,
	HandStage,
	getTeamId,
	getDeclarationTypeName,
} from '@gurian-belote/lib'
import type { InProgressHandState, Declaration } from '@gurian-belote/lib'
import { useGameState } from '../../GameContext'
import SuitIcon from '../../SuitIcon'

interface DeclarationLabelProps {
	decl: Declaration
}

const CARRE_TYPES = new Set([
	DeclarationType.Carre,
	DeclarationType.NinesCarre,
	DeclarationType.JacksCarre,
])

const DeclarationLabel = ({ decl }: DeclarationLabelProps) => {
	const isCarre = CARRE_TYPES.has(decl.type)
	return (
		<>
			{getDeclarationTypeName(decl.type)}
			{decl.highestCard && (
				<span className="inline-flex items-center ml-1">
					{isCarre ? (
						`of ${decl.highestCard.rank}s`
					) : (
						<>
							to {decl.highestCard.rank}
							<SuitIcon
								suit={decl.highestCard.suit}
								iconClassName="w-3 h-3 ml-0.5 inline-block"
							/>
						</>
					)}
				</span>
			)}
		</>
	)
}

interface PlayerDeclarationsProps {
	player: PlayerId
	playerName: string
	decls: Declaration[]
	isLosingTeam: boolean
}

const PlayerDeclarations = ({
	player,
	playerName,
	decls,
	isLosingTeam,
}: PlayerDeclarationsProps) => (
	<>
		{decls.map((decl, i) => (
			<li
				key={`${player}-${i}`}
				className={
					isLosingTeam && decl.type !== DeclarationType.Belote
						? 'line-through opacity-50'
						: ''
				}
			>
				{playerName}: <DeclarationLabel decl={decl} />
				<span className="ml-1">({decl.points} pts)</span>
			</li>
		))}
	</>
)

interface TeamDeclarationListProps {
	declarations: Map<PlayerId, Declaration[]>
	players: Map<PlayerId, string>
	isLosingTeam: boolean
}

const TeamDeclarationList = ({
	declarations,
	players,
	isLosingTeam,
}: TeamDeclarationListProps) => {
	if (declarations.size === 0) return null

	const declarationsArray = Array.from(declarations.entries()).sort()

	return (
		<ul className="text-sm mt-1 opacity-80 list-disc list-inside">
			{declarationsArray.map(([player, decls]) => (
				<PlayerDeclarations
					key={player}
					player={player}
					playerName={players.get(player) ?? `Player ${player}`}
					decls={decls}
					isLosingTeam={isLosingTeam}
				/>
			))}
		</ul>
	)
}

interface DeclarerLabelProps {
	declarerName?: string
}

const DeclarerLabel = ({ declarerName }: DeclarerLabelProps) => (
	<span
		title={declarerName ? `Trump chosen by ${declarerName}` : undefined}
		className="text-xs font-semibold px-2 py-0.5 rounded-full bg-amber-500/20 border border-amber-400/50 text-amber-300"
	>
		Declarers
	</span>
)

interface TeamHandStateProps {
	name: string
	score: number
	declarations: Map<PlayerId, Declaration[]>
	players: Map<PlayerId, string>
	isLosingTeam?: boolean
	isDeclarerTeam?: boolean
	declarerName?: string
}

const TeamHandState = ({
	name,
	score,
	declarations,
	players,
	isLosingTeam = false,
	isDeclarerTeam = false,
	declarerName,
}: TeamHandStateProps) => {
	return (
		<div>
			<h3 className="font-semibold mb-1 border-b border-white/20 pb-1 flex justify-between items-center">
				<span>{name}</span>
				{isDeclarerTeam && (
					<DeclarerLabel declarerName={declarerName} />
				)}
			</h3>
			<p className="text-sm">Points: {score}</p>
			<TeamDeclarationList
				declarations={declarations}
				players={players}
				isLosingTeam={isLosingTeam}
			/>
		</div>
	)
}

const HandStatePanel = () => {
	const { gameState } = useGameState()
	if (!gameState) {
		return null
	}

	if (
		!gameState.gameState.hand ||
		gameState.gameState.hand.state !== HandStage.HandInProgress
	) {
		return null
	}

	const handState = gameState.gameState.hand as InProgressHandState
	const { totals, playerDeclarations, declarationWinner, declarer } =
		handState
	const { players } = gameState.gameState

	const declarerTeam = getTeamId(declarer)
	const declarerName = players.get(declarer) ?? `Player ${declarer}`

	const getDeclarationsForTeam = (
		teamId: TeamId
	): Map<PlayerId, Declaration[]> => {
		const result = new Map<PlayerId, Declaration[]>()
		for (const [playerId, decls] of playerDeclarations.entries()) {
			if (decls.length > 0 && getTeamId(playerId) === teamId) {
				result.set(playerId, decls)
			}
		}
		return result
	}

	return (
		<Panel className="p-6 w-64">
			<h2 className="text-lg font-bold mb-4">Current Hand</h2>
			<div className="space-y-4">
				<TeamHandState
					name="Team 1"
					score={totals.get(TeamId.Team1) || 0}
					declarations={getDeclarationsForTeam(TeamId.Team1)}
					players={players}
					isLosingTeam={declarationWinner === TeamId.Team2}
					isDeclarerTeam={declarerTeam === TeamId.Team1}
					declarerName={declarerName}
				/>
				<TeamHandState
					name="Team 2"
					score={totals.get(TeamId.Team2) || 0}
					declarations={getDeclarationsForTeam(TeamId.Team2)}
					players={players}
					isLosingTeam={declarationWinner === TeamId.Team1}
					isDeclarerTeam={declarerTeam === TeamId.Team2}
					declarerName={declarerName}
				/>
			</div>
		</Panel>
	)
}

export default HandStatePanel
