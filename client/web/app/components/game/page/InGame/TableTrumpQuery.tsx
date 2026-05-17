import Panel from '~/components/Panel'
import CardFace from '../../CardFace'
import type { Card } from '@gurian-belote/lib'

const TableTrumpQuery = (tableTrumpCard: Card, label: string) => {
	return (
		<Panel className="flex flex-col items-center justify-center gap-6 px-8 py-6 text-center">
			<p className="text-lg font-semibold">{label}</p>
			<CardFace card={tableTrumpCard} className="mb-4" />
		</Panel>
	)
}

export default TableTrumpQuery
