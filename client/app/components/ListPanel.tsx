interface ListPanelProps {
	title?: string;
	items: string[];
	footer?: React.ReactNode
	emptyLabel?: string
}

export const ListPanel = ({ title, items, footer, emptyLabel = "No entries yet" }: ListPanelProps) => {
	return (
		<div className="bg-black/60 rounded-xl p-4 shadow-md w-full flex flex-col min-h-[220px]">
			<h2 className="text-white text-lg font-semibold mb-3 text-center">
				{title}
			</h2>

			<div className="flex-1">
				<ul className="overflow-hidden divide-y divide-zinc-300/20">
					{items.length > 0 ? (
						items.map((item, index) => (
							<li
								key={index}
								className="text-white px-4 py-2 hover:bg-zinc-700 transition"
							>
								{item}
							</li>
						))
					) : (
						<div className="text-gray-400 text-sm py-6 text-center">
							{emptyLabel}
						</div>
					)}
				</ul>
			</div>

			{footer && <div className="mt-4">{footer}</div>}
		</div>
	);
}

export default ListPanel;
