import React from "react";
import { Suit, type Card } from "~/client/card";

type CardFaceProps = {
	card: Card;
	hover?: boolean;
};

function getSuitSymbol(suit: Suit): "♠" | "♥" | "♦" | "♣" {
	switch (suit) {
		case Suit.Spades:
			return "♠";
		case Suit.Hearts:
			return "♥";
		case Suit.Diamonds:
			return "♦";
		case Suit.Clubs:
			return "♣";
		default:
			throw new Error("Invalid suit");
	}
}

function getSuitColor(suit: Suit): string {
	switch (suit) {
		case Suit.Spades:
			return "text-black";
		case Suit.Hearts:
			return "text-red-500";
		case Suit.Diamonds:
			return "text-red-500";
		case Suit.Clubs:
			return "text-black";
		default:
			throw new Error("Invalid suit");
	}
}

export const CardFace: React.FC<CardFaceProps> = ({ card, hover = false }) => {
	const hoverAnimationClass = hover ? "transition-transform duration-300 ease-in-out transform hover:scale-105 hover:outline hover:outline-1 hover:outline-[#FFD700]" : "";
	const TopLabel = () => (
		<div className={`absolute top-2 left-2 text-sm font-bold ${getSuitColor(card.suit)}`}>
			{card.rank}
			<br />
			{getSuitSymbol(card.suit)}
		</div>
	);

	const BottomLabel = () => (
		<div
			className={`absolute bottom-2 right-2 text-sm font-bold rotate-180 text-right ${getSuitColor(card.suit)}`}
		>
			{card.rank}
			<br />
			{getSuitSymbol(card.suit)}
		</div>
	);

	const CenterLabel = () => (
		<div className={`text-lg absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 ${getSuitColor(card.suit)}`}>
			{getSuitSymbol(card.suit)}
		</div>
	);


	return (
		<div className={`w-24 h-36 bg-white rounded-xl border-2 border-gray-300 shadow-md p-2 relative ${hoverAnimationClass}`}>
			<TopLabel />
			<CenterLabel />
			<BottomLabel />
		</div>
	);
};

export default CardFace;
