import type { FC } from "react";

interface ActionButtonProps {
	onClick: () => void;
	children?: React.ReactNode;
	disabled?: boolean;
	variant?: 'primary' | 'secondary';
	className?: string;
}

const Button: FC<ActionButtonProps> = ({ 
	onClick, 
	children, 
	disabled = false, 
	variant = 'primary', 
	className = '',
}) => {
	const baseClasses = 'w-full flex items-center justify-center px-4 py-3 font-bold text-sm tracking-wider uppercase rounded-lg focus:outline-none focus:ring-4 disabled:cursor-not-allowed shadow-lg';
	
	const animationClasses = disabled ? '' : 'transition-all duration-300 ease-in-out transform hover:scale-105';

	const variants = {
		primary: 'bg-amber-600 hover:bg-amber-500 text-gray-900 focus:ring-amber-500/50 disabled:bg-gray-700 disabled:text-gray-400 disabled:shadow-none',
		secondary: 'bg-gray-600 hover:bg-gray-500 text-white focus:ring-gray-500/50 disabled:bg-gray-700 disabled:text-gray-400 disabled:shadow-none'
	};

	return (
		<button
			onClick={onClick}
			disabled={disabled}
			className={`${baseClasses} ${animationClasses} ${variants[variant]} ${className}`}
		>
			{children ? children : 'Click Me'}
		</button>
	);
};

export default Button;
