import { LuPlus } from 'react-icons/lu'
import { TeamId } from '@gurian-belote/lib'
import Button from '~/components/Button'
import ListPanel from '~/components/ListPanel'

interface TeamColumnProps {
	teamId: TeamId
	members: string[]
	onJoin: (teamId: TeamId) => void
}

const JoinButton = ({
	teamId,
	onJoin,
}: {
	teamId: TeamId
	onJoin: (teamId: TeamId) => void
}) => (
	<Button onClick={() => onJoin(teamId)} variant="secondary">
		<LuPlus className="inline-block mr-2" />
		Join
	</Button>
)

const TeamColumn = ({ teamId, members, onJoin }: TeamColumnProps) => {
	return (
		<ListPanel
			title={`Team ${teamId === TeamId.Team1 ? 1 : 2}`}
			items={members}
			footer={<JoinButton teamId={teamId} onJoin={onJoin} />}
			emptyLabel="No players yet."
		/>
	)
}

export default TeamColumn
