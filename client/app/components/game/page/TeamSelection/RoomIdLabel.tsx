import { LuCopy } from 'react-icons/lu'

const RoomIdLabel = ({ roomId }: { roomId: string }) => (
	<div className="mb-4 flex items-center justify-center space-x-2">
		<span className="text-sm font-bold text-amber-100/70 tracking-wider">
			{`Room ID: ${roomId}`}
		</span>
		<LuCopy
			className="inline-block cursor-pointer text-gray-400/30 hover:text-amber-100/70 transition-colors duration-200"
			onClick={() => navigator.clipboard.writeText(roomId)}
			size={'1.25rem'}
			title="Copy Room ID"
		/>
	</div>
)

export default RoomIdLabel
