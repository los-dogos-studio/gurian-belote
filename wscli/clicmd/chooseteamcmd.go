package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server"
)

type ChooseTeamCmd struct {
	TeamId game.TeamId
}

type ChooseTeamCmdParser struct{}

func (p *ChooseTeamCmdParser) GetFormat() string {
	return "chooseTeam %d"
}

func (p *ChooseTeamCmdParser) FromInput(input string) (CliCmd, error) {
	cmd := &ChooseTeamCmd{}
	_, err := fmt.Sscanf(input, p.GetFormat(), &cmd.TeamId)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *ChooseTeamCmd) ToWsMsg() (string, error) {
	serverCmd := server.ChooseTeamCmd{
		Command: "chooseTeam",
		TeamId:  cmd.TeamId,
	}
	msg, err := json.Marshal(serverCmd)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
