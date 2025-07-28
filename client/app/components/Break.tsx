interface BreakProps {
	label?: string;
}

export const Break = ({ label }: BreakProps) => {
	return (
		<div className="relative flex py-3 items-center">
			<div className="flex-grow border-t border-gray-700"></div>
			{label && <span className="flex-shrink mx-4 text-gray-500 text-xs uppercase">{label}</span>}
			<div className="flex-grow border-t border-gray-700"></div>
		</div>
	);
}

export default Break;
