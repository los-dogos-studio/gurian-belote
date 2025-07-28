type CardBackgroundProps = {
	children: React.ReactNode;
	className?: string;
};

const CardBackground = ({ children, className = '' }: CardBackgroundProps) => {
	return (
		<div
			className={`w-24 h-36 select-none bg-white rounded-xl border-2 border-gray-300 shadow-md p-2 relative ${className}`}
		>
			{children}
		</div>
	);
};

export default CardBackground;
