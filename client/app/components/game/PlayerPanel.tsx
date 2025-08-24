import { Card, Rank, Suit } from "~/client/card";
import Panel from "../Panel";
import Button from "../Button";
import CardFace from "./CardFace";
import { useGameClient } from "./GameClientContext";
import { useGameState } from "./GameStateContext";
import { GameStage } from "~/client/state/game-state";
import { HandStage, InProgressHandState, TableTrumpSelectionHandState } from "~/client/state/hand";
import { LuClub, LuDiamond, LuHeart, LuSpade } from "react-icons/lu";

interface TableTrumpSelectionPlayerPanelProps {
	cards: Card[];
	disabled: boolean;
}

interface PlayerCardsPanelProps {
	cards: Card[];
	playableCards: Card[];
	disabled: boolean;
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
					<CardFace key={index} card={card} className='brightness-70 cursor-not-allowed pointer-events-none'/>
				))}
			</div>
			{controls}
		</div>
	);

}

const TableTrumpSelectionPlayerPanel = ({ cards, disabled }: TableTrumpSelectionPlayerPanelProps) => {
	const gameClient = useGameClient();

	const TableTrumpSelectionControls = () => (
		<div className="flex flex-col items-center justify-center gap-4">
			<Button
				onClick={() => {
					gameClient.acceptTrump(true);
				}}
				variant="secondary"
				disabled={disabled}
			>
				Accept Table Trump
			</Button>
			<Button
				onClick={() => {
					gameClient.acceptTrump(false);
				}}
				variant="secondary"
				disabled={disabled}
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
	disabled: boolean;
	skippable: boolean;
	iconClassName?: string;
}


const FreeTrumpSelectionPlayerPanel = ({ forbiddenSuit, cards, skippable = true, iconClassName = "", disabled = false }: FreeTrumpSelectionPlayerPanelProps) => {
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

	const TrumpSuitSelectionButton = ({ suit, disabled }: { suit: Suit, disabled: boolean }) => {
		return (
			<Button
				onClick={() => {
					gameClient.selectTrump(suit);
				}}
				disabled={disabled}
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
							<TrumpSuitSelectionButton key={suit} suit={suit} disabled={disabled} />
						)
					))}
				</div>
			</div>
			<Button
				onClick={() => {
					gameClient.selectTrump(null);
				}}
				variant="secondary"
				disabled={!skippable || disabled}
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

const PlayerCardsPanel = ({ cards, playableCards, disabled }: PlayerCardsPanelProps) => {
	const gameClient = useGameClient();

	const onCardClick = (card: Card) => {
		gameClient.playCard(card);
	}

	return (
		<div className="flex justify-center items-center -space-x-4">
			{
				cards.map((card, index) => {
					const notPlayable = !playableCards.some((playableCard) => playableCard.equals(card));
					const cardClassName = (disabled || notPlayable) ? 'brightness-70 cursor-not-allowed pointer-events-none' : '';
					return (
						<CardFace key={index} card={card} onClick={onCardClick} hover className={cardClassName}/>
					)
				})
			}
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
	const currentPlayerId = gameState.gameState.hand?.getCurrentTurn();
	const controlsDisabled = currentPlayerId !== gameState.playerId;

	const skippable = true; // TODO

	let trump = null;
	if (handState === HandStage.TableTrumpSelection || handState === HandStage.FreeTrumpSelection) {
		trump = (hand as TableTrumpSelectionHandState).tableTrumpCard.suit;
	} else {
		trump = (hand as InProgressHandState).trump;
	}
	const cards = gameState.userCards!.sort((a, b) => {
		if (a.suit !== b.suit) {
			const suitOrder = [Suit.Spades, Suit.Hearts, Suit.Diamonds, Suit.Clubs];
			return suitOrder.indexOf(a.suit) - suitOrder.indexOf(b.suit);
		}
		return b.compare(a, trump);
	});


	switch (handState) {
		case HandStage.TableTrumpSelection:
			return <TableTrumpSelectionPlayerPanel disabled={controlsDisabled} cards={cards} />;
		case HandStage.FreeTrumpSelection:
			return <FreeTrumpSelectionPlayerPanel disabled={controlsDisabled} forbiddenSuit={trump} cards={cards} skippable={skippable} />;
		case HandStage.HandInProgress:
			return (
				<PlayerCardsPanel 
					disabled={controlsDisabled} 
					cards={cards} 
					playableCards={hand.getPlayableCards(cards!)}
				/>
			)
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
