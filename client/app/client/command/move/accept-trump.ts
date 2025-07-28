export class AcceptTrumpMove {
	readonly command: string = "acceptTrump";
	accepted: boolean;

	constructor(accepted: boolean) {
		this.accepted = accepted;
	}
}

export default AcceptTrumpMove;
