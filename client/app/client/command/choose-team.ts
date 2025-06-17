import type { TeamId } from "../team-id";

class ChooseTeamCommand {
	readonly command: string = 'chooseTeam';
	teamId: TeamId;

	constructor(teamId: TeamId) {
		this.teamId = teamId;
	}
}

export default ChooseTeamCommand;
