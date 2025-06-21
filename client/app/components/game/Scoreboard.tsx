import Panel from "../Panel";

interface ScoreboardProps {
	scores: Record<string, number>
	className?: string;
};

export const Scoreboard = ({ scores, className = '' }: ScoreboardProps) => (
	<Panel className={"p-8 " + className}>
		<h2 className="text-lg font-bold mb-6">Scores</h2>
		<ul className="space-y-1">
			{Object.entries(scores).map(([playerName, score]) => (
				<li key={playerName} className="flex justify-between">
					<span className="mr-4">{playerName}</span>
					<span>{score}</span>
				</li>
			))}
		</ul>
	</Panel>
);

export default Scoreboard;
