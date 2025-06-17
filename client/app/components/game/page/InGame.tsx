import Panel from "~/components/Panel";
import { useGameState } from "../GameStateContext";
import InputField from "~/components/InputField";
import Button from "~/components/Button";
import { useState } from "react";
import { useGameClient } from "../GameClientContext";
import { Rank, Suit } from "~/client/card";

export const InGame = () => {
	const { gameState } = useGameState();
	const gameClient = useGameClient();
	const [inputtedCard, setInputtedCard] = useState("");
	const [inputtedTrump, setInputtedTrump] = useState("");

	const playCard = () => {
		if (!gameState) {
			throw new Error("Game state is not available");
		}

		if (!inputtedCard) {
			throw new Error("No card inputted");
		}

		const cardParts = inputtedCard.split(" ");
		if (cardParts.length < 2) {
			throw new Error("Invalid card format");
		}
		const suit = cardParts[0];
		const rank = cardParts[1];
		const card = {
			suit: suit as Suit,
			rank: rank as Rank
		};

		console.log(card);
		gameClient.playCard(card);
	}

	return (
		<Panel>
			<h1>
				In Game
			</h1>
			<div style={{ marginBottom: "20px" }}>
				{JSON.stringify(gameState, null, "\t")}
			</div>
			<InputField label={"Enter Card"} value={inputtedCard} onChange={setInputtedCard} placeholder={"E.g. K8"} isAlphaNumeric={false} />
			<Button
				onClick={playCard}
				className="my-2"
				variant="secondary"
			>
				Play Card
			</Button>
			<Button
				onClick={() => gameClient.acceptTrump(true)}
				className="my-2"
				variant="secondary"
			>
				Accept Trump
			</Button>
			<Button
				onClick={() => gameClient.acceptTrump(false)}
				className="my-2"
				variant="secondary"
			>
				Reject Trump
			</Button>
			<InputField label={"Select Trump"} value={inputtedTrump} onChange={setInputtedTrump} placeholder={"E.g. Hearts"} />
			<Button
				onClick={() => gameClient.selectTrump(inputtedTrump.length > 0 ? inputtedTrump as Rank : null)}
				className="my-2"
				disabled={!inputtedTrump}
				variant="secondary"
			>
				Select Trump
			</Button>
		</Panel>
	);
}

export default InGame;
