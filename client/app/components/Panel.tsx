export const Panel: React.FC<{
	children: React.ReactNode;
	className?: string;
}> = ({ children, className }) => {
	return (
		<div
			className={`bg-black/70 backdrop-blur-sm rounded-2xl shadow-2xl p-8 max-w-md w-full ${className}`}
		>
			{children}
		</div>
	);
};

export default Panel;
