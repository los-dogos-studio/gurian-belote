import { LuUser } from 'react-icons/lu'

const PlayerIcon = ({
	label,
	className = '',
}: {
	label: string
	className?: string
}) => {
	return (
		<div className="flex flex-col items-center justify-center">
			<div
				className={`w-12 h-12 bg-gray-800 rounded-full mx-auto mb-2 flex items-center justify-center ${className}`}
			>
				<LuUser className="w-6 h-6 text-white" />
			</div>
			<p className="text-s">{label}</p>
		</div>
	)
}

export default PlayerIcon
