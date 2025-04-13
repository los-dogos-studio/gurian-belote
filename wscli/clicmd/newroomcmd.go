package clicmd

import (
	"encoding/json"

	"github.com/los-dogos-studio/gurian-belote/server"
)

type NewRoomCmd struct{}

type NewRoomCmdParser struct{}

func (p *NewRoomCmdParser) GetFormat() string {
	return "newRoom"
}

func (p *NewRoomCmdParser) FromInput(input string) (CliCmd, error) {
	if input != p.GetFormat() {
		return nil, ErrInvalidCmd
	}
	return &NewRoomCmd{}, nil
}

func (cmd *NewRoomCmd) ToWsMsg() (string, error) {
	serverCmd := server.Cmd{
		Command: "newRoom",
	}
	msg, err := json.Marshal(serverCmd)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
