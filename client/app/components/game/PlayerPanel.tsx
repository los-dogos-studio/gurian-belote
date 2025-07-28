import { Suit, type Card } from "~/client/card";
import Panel from "../Panel";
import Button from "../Button";
import CardFace from "./CardFace";
import { useGameClient } from "./GameClientContext";
import { useGameState } from "./GameStateContext";
import { GameStage } from "~/client/state/game-state";
import { HandStage, TableTrumpSelectionHandState } from "~/client/state/hand";
import { LuClub, LuDiamond, LuHeart, LuSpade } from "react-icons/lu";

interface TableTrumpSelectionPlayerPanelProps {
	cards: Card[];
}

interface PlayerCardsPanelProps {
	cards: Card[];
}

interface TrumpSelectionPlayerPanelProps {
	cards: Card[];
	controls: React.ReactNode;
}

const TrumpSelectionPlayerPanel = ({ cards, controls }: TrumpSelectionPlayerPanelProps) => {
	return (
		<div className="flex items-center justify-center gap-8">
			<div className="flex justify-center items-center -space-x-4">
				{cards.map((card, index) => (
					<CardFace key={index} card={card} hover />
				))}
			</div>
			{controls}
		</div>
	);

}

const TableTrumpSelectionPlayerPanel = ({ cards }: TableTrumpSelectionPlayerPanelProps) => {
	const gameClient = useGameClient();

	const TableTrumpSelectionControls = () => (
		<div className="flex flex-col items-center justify-center gap-4">
			<Button
				onClick={() => {
					gameClient.acceptTrump(true);
				}}
				variant="secondary"
			>
				Accept Table Trump
			</Button>
			<Button
				onClick={() => {
					gameClient.acceptTrump(false);
				}}
				variant="secondary"
			>
				Decline Table Trump
			</Button>
		</div>
	);


	return (
		<TrumpSelectionPlayerPanel
			cards={cards}
			controls={<TableTrumpSelectionControls />}
		/>
	);
}

interface FreeTrumpSelectionPlayerPanelProps {
	forbiddenSuit: Suit;
	cards: Card[];
	skippable: boolean;
	iconClassName?: string;
}


const FreeTrumpSelectionPlayerPanel = ({ forbiddenSuit, cards, skippable = true, iconClassName = "" }: FreeTrumpSelectionPlayerPanelProps) => {
	const gameClient = useGameClient();

	const SuitIcon = ({ suit }: { suit: Suit }) => {
		switch (suit) {
			case Suit.Spades:
				return <LuSpade className={iconClassName} />;
			case Suit.Hearts:
				return <LuHeart className={iconClassName} />;
			case Suit.Diamonds:
				return <LuDiamond className={iconClassName} />;
			case Suit.Clubs:
				return <LuClub className={iconClassName} />;
			default:
				throw new Error("Invalid suit");
		}
	}

	const TrumpSuitSelectionButton = ({ suit }: { suit: Suit }) => {
		return (
			<Button
				onClick={() => {
					gameClient.selectTrump(suit);
				}}
				variant="secondary"
			>
				{<SuitIcon suit={suit} />}
			</Button>
		);
	}

	const FreeTrumpSelectionControls = () => (
		<div className="flex flex-col items-center justify-center gap-4">
			<div>
				<div className="flex justify-center items-center gap-2">
					{Object.values(Suit).map((suit) => (
						suit !== forbiddenSuit && (
							<TrumpSuitSelectionButton key={suit} suit={suit} />
						)
					))}
				</div>
			</div>
			<Button
				onClick={() => {
					gameClient.selectTrump(null);
				}}
				variant="secondary"
				disabled={!skippable}
			>
				Skip
			</Button>
		</div>
	);

	return (
		<TrumpSelectionPlayerPanel
			cards={cards}
			controls={<FreeTrumpSelectionControls />}
		/>
	);
}

const PlayerCardsPanel = ({ cards }: PlayerCardsPanelProps) => {
	const gameClient = useGameClient();

	const onCardClick = (card: Card) => {
		gameClient.playCard(card);
	}

	return (
		<div className="flex justify-center items-center -space-x-4">
			{cards.map((card, index) => (
				<CardFace key={index} card={card} onClick={onCardClick} hover />
			))}
		</div>
	);
}

const PlayerPanelContent = () => {
	const { gameState } = useGameState();
	if (
		!gameState ||
		!gameState.gameState ||
		!gameState.gameState.hand ||
		!gameState.gameState.gameState ||
		gameState.gameState.gameState !== GameStage.GameInProgress
	) {
		return <div>Invalid game stage...</div>;
	}

	const hand = gameState.gameState.hand;
	const handState = hand.state;
	const cards = gameState.userCards;

	const skippable = true; // TODO

	switch (handState) {
		case HandStage.TableTrumpSelection:
			return <TableTrumpSelectionPlayerPanel cards={cards!} />;
		case HandStage.FreeTrumpSelection:
			return <FreeTrumpSelectionPlayerPanel forbiddenSuit={(hand as TableTrumpSelectionHandState).tableTrumpCard.suit} cards={cards!} skippable={skippable} />;
		case HandStage.HandInProgress:
			return <PlayerCardsPanel cards={cards!} />;
		default:
			return <div>Invalid hand state...</div>;
	}
}

export const PlayerPanel = ({ className = '' }: { className?: string }) => {
	return (
		<Panel className={`flex justify-center p-4 mb-2 gap-4 ${className}`}>
			<PlayerPanelContent />
		</Panel>
	);
}

export default PlayerPanel;
