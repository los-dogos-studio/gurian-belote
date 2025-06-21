import { useGameState } from "../GameStateContext";
import { useGameClient } from "../GameClientContext";
import Scoreboard from "../Scoreboard";
import { Rank, Suit, type Card } from "~/client/card";
import PlayerPanel from "../PlayerPanel";
import { LuUser } from "react-icons/lu";
import CardFace from "../CardFace";
import { getNextPlayerId, type PlayerId } from "~/client/player-id";

const scores = {
	"You": 0,
	"Player 2": 0,
	"Player 3": 0,
	"Player 4": 0,
};

interface PlayedCardsProps {
	bottom: Card | undefined;
	left: Card | undefined;
	top: Card | undefined;
	right: Card | undefined;
}

const PlayedCards = ({ bottom, left, top, right }: PlayedCardsProps) => {
	const CardSlot = ({ card, className = '' }: { card: Card | undefined, className?: string }) => {
		return (
			<div className={className}>
				{card && <CardFace card={card} />}
			</div>
		);
	}

	// rethink?
	// TODO: add z
	return (
		<div className="relative w-full h-full">
			<CardSlot card={bottom} className='absolute bottom-2 left-1/2 transform -translate-x-1/2' />
			<CardSlot card={left} className='absolute left-2 top-1/2 transform -translate-y-1/2' />
			<CardSlot card={top} className='absolute top-2 left-1/2 transform -translate-x-1/2' />
			<CardSlot card={right} className='absolute right-2 top-1/2 transform -translate-y-1/2' />
		</div>
	)
}

const playedCards: PlayedCardsProps = {
	"bottom": { rank: Rank.Ace, suit: Suit.Hearts },
	"left": { rank: Rank.King, suit: Suit.Spades },
	"top": { rank: Rank.Queen, suit: Suit.Diamonds },
	"right": { rank: Rank.Jack, suit: Suit.Clubs },
};

export const InGame = () => {
	const { gameState } = useGameState();
	const gameClient = useGameClient();

	if (!gameState) {
		return <div className="text-white">Waiting for game...</div>;
	}

	const PlayerIcon = ({ label }: { label: string }) => {
		return (
			<div className="flex flex-col items-center justify-center">
				<div className="w-12 h-12 bg-gray-800 rounded-full mx-auto mb-2 flex items-center justify-center">
					<LuUser className="w-6 h-6 text-white" />
				</div>
				<p className="text-s">{label}</p>
			</div>
		);
	}

	const leftPlayerId: PlayerId = getNextPlayerId(gameState.playerId);
	const topPlayerId: PlayerId = getNextPlayerId(leftPlayerId);
	const rightPlayerId: PlayerId = getNextPlayerId(topPlayerId);

	const leftPlayerName = gameState.gameState.players.get(leftPlayerId) ?? `Player ${leftPlayerId}`;
	const topPlayerName = gameState.gameState.players.get(topPlayerId) ?? `Player ${topPlayerId}`;
	const rightPlayerName = gameState.gameState.players.get(rightPlayerId) ?? `Player ${rightPlayerId}`;


	return (
		<div className="h-full w-full relative gap-4 p-4 text-white">
			<div className="absolute top-2 right-2 p-3">
				<Scoreboard scores={scores} className="h-min" />
			</div>

			<div className="absolute top-1/2 left-2 transform -translate-y-1/2 p-3">
				<PlayerIcon label={leftPlayerName} />
			</div>

			<div className="absolute top-2 left-1/2 transform -translate-x-1/2 p-3">
				<PlayerIcon label={topPlayerName} />
			</div>

			<div className="absolute top-1/2 right-2 transform -translate-y-1/2 p-3">
				<PlayerIcon label={rightPlayerName} />
			</div>

			<div className="absolute bottom-2 left-1/2 transform -translate-x-1/2 p-3">
				<PlayerPanel />
			</div>

			<div className="inline-block absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
				<PlayedCards
					bottom={playedCards['bottom']}
					left={playedCards['left']}
					top={playedCards['top']}
					right={playedCards['right']}
				/>
			</div>
		</div>
	);
}

export default InGame;
