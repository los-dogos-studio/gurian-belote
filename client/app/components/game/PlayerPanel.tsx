import { Rank, Suit, type Card } from "~/client/card";
import Panel from "../Panel";
import Button from "../Button";
import CardFace from "./CardFace";
import { useGameClient } from "./GameClientContext";
import { useGameState } from "./GameStateContext";
import { GameStage } from "~/client/state/game-state";
import { HandState } from "~/client/state/hand";
import { LuClub, LuDiamond, LuHeart, LuSpade } from "react-icons/lu";

interface PlayerCardsPanelProps {
	cards: Card[];
}

const myCards = [
	{ suit: Suit.Spades, rank: Rank.Ace },
	{ suit: Suit.Hearts, rank: Rank.King },
	{ suit: Suit.Diamonds, rank: Rank.Queen },
	{ suit: Suit.Clubs, rank: Rank.Jack },
	{ suit: Suit.Spades, rank: Rank.Ten },
];

const TableTrumpSelectionPlayerPanel = () => {
	return (
		<div className="flex flex-col items-center justify-center gap-4">
			<Button
				onClick={() => {
					// Handle accepting the table trump
					console.log("Table Trump Accepted");
				}}
				variant="secondary"
			>
				Accept Table Trump
			</Button>
			<Button
				onClick={() => {
					// Handle declining the table trump
					console.log("Table Trump Declined");
				}}
				variant="secondary"
			>
				Decline Table Table
			</Button>
		</div>
	);
}

interface FreeTrumpSelectionPlayerPanelProps {
	forbiddenSuit: Suit;
	skippable: boolean;
	iconClassName?: string;
}


const FreeTrumpSelectionPlayerPanel = ({ forbiddenSuit, skippable = true, iconClassName = "" }: FreeTrumpSelectionPlayerPanelProps) => {
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
					// Handle selecting the trump suit
					console.log(`Trump Suit Selected: ${suit}`);
				}}
				variant="secondary"
			>
				{<SuitIcon suit={suit} />}
			</Button>
		);
	}

	return (
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
			{skippable && (
				<Button
					onClick={() => {
						// Handle skipping the free trump selection
						console.log("Free Trump Selection Skipped");
					}}
					variant="secondary"
				>
					Skip
				</Button>
			)}
		</div>
	);
}

const PlayerCardsPanel = ({ cards }: PlayerCardsPanelProps) => {
	return (
		<div className="flex justify-center items-center -space-x-4">
			{cards.map((card, index) => (
				<CardFace key={index} card={card} hover />
			))}
		</div>
	);
}

const PlayerPanelContent = () => {
	// const { gameState } = useGameState();
	// if (
	// 	!gameState ||
	// 	!gameState.gameState ||
	// 	!gameState.gameState.hand ||
	// 	!gameState.gameState.gameState ||
	// 	gameState.gameState.gameState !== GameStage.GameInProgress
	// ) {
	// 	return <div>Invalid game stage...</div>;
	// }

	// const handState = gameState.gameState.hand.state;
	// const cards = gameState.userCards;
	const handState = HandState.HandInProgress;

	switch (handState) {
		case HandState.TableTrumpSelection:
			return <TableTrumpSelectionPlayerPanel />;
		case HandState.FreeTrumpSelection:
			return <FreeTrumpSelectionPlayerPanel forbiddenSuit={Suit.Clubs} />;
		case HandState.HandInProgress:
			return <PlayerCardsPanel cards={myCards} />;
		case HandState.HandFinished:
			return <PlayerCardsPanel cards={myCards} />;
		default:
			return <div>Invalid hand state...</div>;
	}
}

export const PlayerPanel = () => {
	return (
		<Panel className="flex justify-center p-4 mb-2 gap-4">
			<PlayerPanelContent />
		</Panel>
	);
}

export default PlayerPanel;
