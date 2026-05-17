import Panel from '~/components/Panel'
import { getSuitSymbol } from '../../card-utils'
import type { Suit } from '@gurian-belote/lib'

const TrumpDisplay = ({ suit }: { suit: Suit }) => {
	return (
		<Panel className="w-18 h-18 flex flex-col items-center justify-center gap-3 text-center">
			<span
				className={`text-[3em] w-full h-full flex items-center justify-center text-center text-white/80`}
			>
				{getSuitSymbol(suit)}
			</span>
		</Panel>
	)
}

export default TrumpDisplay
