import { useGameState } from "../GameStateContext";
import { useGameClient } from "../GameClientContext";
import Scoreboard from "../Scoreboard";
import { Rank, Suit, type Card } from "~/client/card";
import PlayerPanel from "../PlayerPanel";
import { LuUser } from "react-icons/lu";
import CardFace from "../CardFace";

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

	return (
		<div className="relative w-min h-full gap-4">
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

	return (
		<div className="grid grid-cols-3 grid-rows-3 gap-4 h-screen w-screen p-4 text-white">
			<div className="col-start-3 row-start-1 flex justify-end" >
				<Scoreboard scores={scores} className="h-min" />
			</div>

			<div className="col-start-1 row-start-2 flex justify-start items-center">
				<PlayerIcon label="Player 2" />
			</div>

			<div className="col-start-3 row-start-2 flex justify-end items-center">
				<PlayerIcon label="Player 3" />
			</div>

			<div className="col-start-2 row-start-1 flex flex-col justify-start items-center">
				<PlayerIcon label="Player 4" />
			</div>

			<div className="col-start-2 row-start-2 flex justify-center items-center">
				<PlayedCards
					bottom={playedCards['bottom']}
					left={playedCards['left']}
					top={playedCards['top']}
					right={playedCards['right']}
				/>
			</div>

			<div className="col-start-2 row-start-3 flex flex-col justify-end items-center">
				<PlayerPanel />
			</div>
		</div>
	);
}

export default InGame;
