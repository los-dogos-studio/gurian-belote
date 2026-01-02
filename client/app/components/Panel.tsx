export const Panel = ({
	children,
	className,
}: {
	children: React.ReactNode
	className?: string
}) => {
	return (
		<div
			className={`bg-black/70 backdrop-blur-sm rounded-2xl shadow-2xl p-2 ${className}`}
		>
			{children}
		</div>
	)
}

export default Panel
