import type { Card } from '@gurian-belote/lib'
import { useGameClient } from '../GameContext'
import Button from '~/components/Button'
import TrumpSelectionPlayerPanel from './TrumpSelectionPlayerPanel'

interface TableTrumpSelectionPlayerPanelProps {
	cards: Card[]
	disabled: boolean
}

const TableTrumpSelectionControls = ({ disabled }: { disabled: boolean }) => {
	const gameClient = useGameClient()

	return (
		<div className="flex flex-col items-center justify-center gap-4">
			<Button
				onClick={() => {
					gameClient.acceptTrump(true)
				}}
				variant="secondary"
				disabled={disabled}
			>
				Accept Table Trump
			</Button>
			<Button
				onClick={() => {
					gameClient.acceptTrump(false)
				}}
				variant="secondary"
				disabled={disabled}
			>
				Decline Table Trump
			</Button>
		</div>
	)
}

const TableTrumpSelectionPlayerPanel = ({
	cards,
	disabled,
}: TableTrumpSelectionPlayerPanelProps) => (
	<TrumpSelectionPlayerPanel
		cards={cards}
		controls={<TableTrumpSelectionControls disabled={disabled} />}
	/>
)

export default TableTrumpSelectionPlayerPanel
