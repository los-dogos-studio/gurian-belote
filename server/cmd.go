package server

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/los-dogos-studio/gurian-belote/game"
	"github.com/los-dogos-studio/gurian-belote/server/internal/serverroom"
	"github.com/los-dogos-studio/gurian-belote/server/internal/serverroom/gamecmd"
)

type CmdName string

type Cmd struct {
	Command CmdName
}

type JoinRoomCmd struct {
	Command CmdName
	RoomId  string
}

type ChooseTeamCmd struct {
	Command CmdName
	TeamId  game.TeamId
}

type PlayTurnCmd struct {
	Command CmdName
	Move    json.RawMessage
}

const (
	NewRoom    CmdName = "newRoom"
	JoinRoom   CmdName = "joinRoom"
	ChooseTeam CmdName = "chooseTeam"
	StartGame  CmdName = "startGame"
	PlayTurn   CmdName = "playTurn"
)

var (
	ErrInvalidCmdType   = errors.New("invalid command type")
	ErrInvalidTeamId    = errors.New("invalid team id")
	ErrInvalidCmdParams = errors.New("invalid command parameters")
	ErrUserHasRoom      = errors.New("user already has a room")
	ErrUserNotInRoom    = errors.New("user not in room")
	ErrRoomNotFound     = errors.New("room not found")
)

// TODO: Refactor
func (s *Server) handleCommand(user *User, content json.RawMessage) error {
	cmd := Cmd{}
	err := json.Unmarshal(content, &cmd)
	if err != nil {
		return err
	}

	switch cmd.Command {
	case NewRoom:
		if user.RoomId != "" {
			return ErrUserHasRoom
		}

		userRoom := serverroom.NewRoom(strconv.Itoa(s.getNextRoomId()))
		s.rooms[userRoom.Id] = userRoom

		err = userRoom.Join(user.Id)
		if err != nil {
			delete(s.rooms, userRoom.Id)
			return err
		}

		user.RoomId = userRoom.Id
		// user.ws.WriteMessage(1, []byte(userRoom.Id))
		return nil

	case JoinRoom:
		if user.RoomId != "" {
			return ErrUserHasRoom
		}

		joinRoomCommand := JoinRoomCmd{}

		err := json.Unmarshal(content, &joinRoomCommand)
		if err != nil {
			return err
		}

		userRoom := s.rooms[joinRoomCommand.RoomId]
		if userRoom == nil {
			return ErrRoomNotFound
		}

		err = userRoom.Join(user.Id)
		if err != nil {
			return err
		}

		user.RoomId = userRoom.Id
		return nil

	case StartGame:
		if user.RoomId == "" {
			return ErrUserNotInRoom
		}

		userRoom := s.rooms[user.RoomId]
		return userRoom.StartGame()

	case ChooseTeam:
		chooseTeamCmd := ChooseTeamCmd{}
		err := json.Unmarshal(content, &chooseTeamCmd)
		if err != nil {
			return err
		}

		if chooseTeamCmd.TeamId != game.Team1 && chooseTeamCmd.TeamId != game.Team2 {
			// TODO: check if all conditions are actually needed
			return ErrInvalidTeamId
		}

		if user.RoomId == "" {
			return ErrUserNotInRoom
		}

		userRoom := s.rooms[user.RoomId]
		return userRoom.ChooseTeam(user.Id, chooseTeamCmd.TeamId)

	case PlayTurn:
		playTurnCmd := PlayTurnCmd{}
		err := json.Unmarshal(content, &playTurnCmd)
		if err != nil {
			return err
		}

		if playTurnCmd.Move == nil {
			return ErrInvalidCmdParams
		}

		roomCmd, err := gamecmd.NewGameCmdFromJson(playTurnCmd.Move)
		if err != nil {
			return err
		}

		if user.RoomId == "" {
			return ErrUserNotInRoom
		}

		userRoom := s.rooms[user.RoomId]
		return userRoom.PlayTurn(user.Id, roomCmd)

	default:
		return ErrInvalidCmdType
	}
}

func (s *Server) getNextRoomId() int {
	roomId := s.roomIdCounter
	s.roomIdCounter++
	return roomId
}
