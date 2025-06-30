import React from "react";
import { type Card } from "~/client/card";
import CardBackground from "./CardBackground";
import { getSuitColor, getSuitSymbol } from "./card-utils";

type CardFaceProps = {
	card: Card;
	hover?: boolean;
	onClick?: (card: Card) => void;
	className?: string;
};

const CardFace: React.FC<CardFaceProps> = ({ card, hover = false, onClick = () => { }, className = '' }) => {
	const hoverAnimationClass = hover ? "transition-transform duration-300 ease-in-out transform hover:scale-105 hover:outline hover:outline-1 hover:outline-[#FFD700]" : "";
	const TopLabel = () => (
		<div className={`text-sm font-bold absolute top-2 left-2 ${getSuitColor(card.suit)}`}>
			{card.rank}
			< br />
			{getSuitSymbol(card.suit)}
		</div >
	);

	const BottomLabel = () => (
		<div
			className={`text-sm font-bold rotate-180 text-right absolute bottom-2 right-2 ${getSuitColor(card.suit)}`}
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

	// FIXME: select-none doesn't work in Safari
	return (
		<div
			onClick={() => { onClick(card) }}
		>
			<CardBackground className={`${hoverAnimationClass} ${className}`}>
				<TopLabel />
				<CenterLabel />
				<BottomLabel />
			</CardBackground>
		</div>
	);
};

export default CardFace;
