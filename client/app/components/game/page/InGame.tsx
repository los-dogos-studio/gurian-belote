import { useGameState } from "../GameStateContext";
import Scoreboard from "../Scoreboard";
import { Suit, type Card } from "~/client/card";
import PlayerPanel from "../PlayerPanel";
import { LuUser } from "react-icons/lu";
import CardFace from "../CardFace";
import { getNextPlayerId, type PlayerId } from "~/client/player-id";
import { FreeTrumpSelectionHandState, HandStage, InProgressHandState, TableTrumpSelectionHandState } from "~/client/state/hand";
import { TeamId } from "~/client/team-id";
import Panel from "~/components/Panel";
import { getSuitSymbol } from "../card-utils";

interface TrickProps {
	bottom: Card | undefined;
	left: Card | undefined;
	top: Card | undefined;
	right: Card | undefined;
}

const Trick = ({ bottom, left, top, right }: TrickProps) => {
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

export const InGame = () => {
	const { gameState } = useGameState();

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

	const TableTrumpQuery = (tableTrumpCard: Card, label: string) => {
		return (
			<Panel className="flex flex-col items-center justify-center gap-6 px-8 py-6 text-center">
				<p className="text-lg font-semibold">{label}</p>
				<CardFace card={tableTrumpCard} className="mb-4" />
			</Panel>
		);
	}

	const TrumpDisplay = ({ suit }: { suit: Suit }) => {
		return (
			<Panel className="w-18 h-18 flex flex-col items-center justify-center gap-3 text-center">
				<span className={`text-[3em] w-full h-full flex items-center justify-center text-center text-white/80`}>
					{getSuitSymbol(suit)}
				</span>
			</Panel>
		);
	}

	const PlayArea = () => {
		if (!gameState.gameState.hand) {
			return <div />;
		}

		switch (gameState.gameState.hand.state) {
			case HandStage.TableTrumpSelection:
				return TableTrumpQuery((gameState.gameState.hand as TableTrumpSelectionHandState).tableTrumpCard, "Accept Table Trump?");
			case HandStage.FreeTrumpSelection:
				return TableTrumpQuery((gameState.gameState.hand as FreeTrumpSelectionHandState).tableTrumpCard, "Select Table Trump");

			case HandStage.HandInProgress:
				const inProgressHand = gameState.gameState.hand as InProgressHandState;
				return (
					<Trick
						bottom={inProgressHand.trick.playedCards.get(gameState.playerId)}
						left={inProgressHand.trick.playedCards.get(leftPlayerId)}
						top={inProgressHand.trick.playedCards.get(topPlayerId)}
						right={inProgressHand.trick.playedCards.get(rightPlayerId)}
					/>
				);

			default:
				return (
					<div />
				);
		}
	}

	const leftPlayerId: PlayerId = getNextPlayerId(gameState.playerId);
	const topPlayerId: PlayerId = getNextPlayerId(leftPlayerId);
	const rightPlayerId: PlayerId = getNextPlayerId(topPlayerId);

	const leftPlayerName = gameState.gameState.players.get(leftPlayerId) ?? `Player ${leftPlayerId}`;
	const topPlayerName = gameState.gameState.players.get(topPlayerId) ?? `Player ${topPlayerId}`;
	const rightPlayerName = gameState.gameState.players.get(rightPlayerId) ?? `Player ${rightPlayerId}`;

	const scores = {
		"Team 1": gameState.gameState.scores.get(TeamId.Team1) ?? -1,
		"Team 2": gameState.gameState.scores.get(TeamId.Team2) ?? -1,
	};

	return (
		<div className="h-full w-full relative gap-4 p-4 text-white">
			<div className="absolute top-2 right-2 p-3 flex gap-4">
				{gameState.gameState.hand && gameState.gameState.hand.state === HandStage.HandInProgress &&
					<TrumpDisplay suit={(gameState.gameState.hand as InProgressHandState).trump} />
				}
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
				<PlayArea />
			</div>
		</div>
	);
}

export default InGame;
