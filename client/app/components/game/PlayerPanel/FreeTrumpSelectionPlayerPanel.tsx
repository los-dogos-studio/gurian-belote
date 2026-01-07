import Button from '~/components/Button'
import TrumpSelectionPlayerPanel from './TrumpSelectionPlayerPanel'
import { Card, Suit } from '~/client/card'
import SuitIcon from '../SuitIcon'
import { useGameClient } from '../GameContext'

interface TrumpSuitSelectionButtonProps {
	suit: Suit
	disabled: boolean
	onSelect: () => void
}

interface FreeTrumpSelectionControlsProps {
	forbiddenSuit: Suit
	disabled: boolean
	skippable: boolean
	onSelect: (suit: Suit | null) => void
}

interface FreeTrumpSelectionPlayerPanelProps {
	forbiddenSuit: Suit
	cards: Card[]
	disabled: boolean
	skippable: boolean
}

interface TrumpSelectionSkipButtonProps {
	disabled: boolean
	onClick: () => void
}

const TrumpSuitSelectionButton = ({
	suit,
	disabled,
	onSelect,
}: TrumpSuitSelectionButtonProps) => (
	<Button onClick={onSelect} disabled={disabled} variant="secondary">
		{<SuitIcon suit={suit} />}
	</Button>
)

const TrumpSelectionSkipButton = ({
	disabled,
	onClick,
}: TrumpSelectionSkipButtonProps) => (
	<Button onClick={onClick} variant="secondary" disabled={disabled}>
		Skip
	</Button>
)

const FreeTrumpSelectionControls = ({
	forbiddenSuit,
	disabled,
	skippable,
	onSelect,
}: FreeTrumpSelectionControlsProps) => (
	<div className="flex flex-col items-center justify-center gap-4">
		<div>
			<div className="flex justify-center items-center gap-2">
				{Object.values(Suit).map(
					(suit) =>
						suit !== forbiddenSuit && (
							<TrumpSuitSelectionButton
								key={suit}
								suit={suit}
								disabled={disabled}
								onSelect={() => {
									onSelect(suit)
								}}
							/>
						)
				)}
			</div>
		</div>
		<TrumpSelectionSkipButton
			disabled={!skippable || disabled}
			onClick={() => onSelect(null)}
		/>
	</div>
)

const FreeTrumpSelectionPlayerPanel = ({
	forbiddenSuit,
	cards,
	skippable = true,
	disabled = false,
}: FreeTrumpSelectionPlayerPanelProps) => {
	const gameClient = useGameClient()

	return (
		<TrumpSelectionPlayerPanel
			cards={cards}
			controls={
				<FreeTrumpSelectionControls
					forbiddenSuit={forbiddenSuit}
					disabled={disabled}
					skippable={skippable}
					onSelect={function (suit: Suit | null): void {
						gameClient.selectTrump(suit)
					}}
				/>
			}
		/>
	)
}

export default FreeTrumpSelectionPlayerPanel
