interface SkipDeclarationsToggleProps {
	active: boolean
	disabled?: boolean
	onToggle: () => void
}

const SkipDeclarationsToggle = ({
	active,
	disabled = false,
	onToggle,
}: SkipDeclarationsToggleProps) => {
	return (
		<button
			id="skip-declarations-toggle"
			onClick={onToggle}
			disabled={disabled}
			title={
				active
					? 'Declarations will be skipped for your next played card'
					: 'Click to skip declarations for your next played card'
			}
			className={[
				'flex items-center gap-2 px-3 py-1.5 rounded-full text-xs font-semibold tracking-wide border transition-all duration-200',
				disabled
					? 'opacity-40 cursor-not-allowed'
					: 'cursor-pointer hover:scale-105',
				active
				? 'bg-amber-500/20 border-amber-400 text-amber-300 shadow shadow-amber-400/40 hover:shadow-md hover:shadow-amber-400/60 hover:border-amber-300'
				: 'bg-gray-700/60 border-gray-500 text-gray-400 hover:border-gray-400 hover:text-gray-300',
			].join(' ')}
		>
			<span
				className={[
					'w-2 h-2 rounded-full transition-colors duration-200',
					active ? 'bg-amber-400' : 'bg-gray-500',
				].join(' ')}
			/>
			Skip declarations
		</button>
	)
}

export default SkipDeclarationsToggle
