package clicmd

import (
	"encoding/json"
	"fmt"

	"github.com/los-dogos-studio/gurian-belote/server"
)

type JoinRoomCmd struct {
	RoomId string
}

type JoinRoomCmdParser struct{}

func (p *JoinRoomCmdParser) GetFormat() string {
	return "joinRoom %s"
}

func (p *JoinRoomCmdParser) FromInput(input string) (CliCmd, error) {
	var roomId string
	_, err := fmt.Sscanf(input, p.GetFormat(), &roomId)
	if err != nil {
		return nil, err
	}
	return &JoinRoomCmd{RoomId: roomId}, nil
}

func (cmd *JoinRoomCmd) ToWsMsg() (string, error) {
	serverCmd := server.JoinRoomCmd{
		Command: "joinRoom",
		RoomId:  cmd.RoomId,
	}
	msg, err := json.Marshal(serverCmd)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
