import type AcceptTrumpMove from "./accept-trump";
import type PlayCardMove from "./play-card";
import type SelectTrumpMove from "./select-trump";

export type GameMove =
	| AcceptTrumpMove
	| SelectTrumpMove
	| PlayCardMove;
