export const Background = ({ children }: { children: React.ReactNode }) => (
	<div className="min-h-screen bg-gradient-to-br from-[#0d1b2a] via-[#1b263b] to-black flex items-center justify-center p-4">
		 {/* <div className="min-h-screen bg-gradient-to-br from-[#2b0000] via-[#400000] to-black flex items-center justify-center p-4"> */}
		{children}
	</div>
);

export default Background;
