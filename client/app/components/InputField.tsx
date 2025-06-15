import type { FC } from 'react';

interface InputFieldProps {
	label: string;
	value: string;
	onChange: (value: string) => void;
	placeholder: string;
	error?: string;
	disabled?: boolean;
}

const InputField: FC<InputFieldProps> = ({ label, value, onChange, placeholder, error, disabled = false }) => {
	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const inputValue = e.target.value;
		if (/^[a-zA-Z0-9]*$/.test(inputValue)) {
			onChange(inputValue);
		}
	};

	return (
		<div className="mb-4">
			<label className="block text-sm font-bold text-amber-100/70 mb-2 tracking-wider">
				{label}
			</label>
			<input
				type="text"
				value={value}
				onChange={handleChange}
				placeholder={placeholder}
				disabled={disabled}
				className={`w-full px-4 py-3 bg-gray-800/50 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition duration-200 text-white placeholder-gray-500 disabled:opacity-50 shadow-inner`}
			/>
			{error && <p className="text-red-400 text-xs mt-1">{error}</p>}
		</div>
	);
};

export default InputField;
