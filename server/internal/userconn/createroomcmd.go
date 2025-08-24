package userconn

type CreateRoomCmd struct {
}

func NewCreateRoomCmd(msg []byte) (Cmd, error) {
	return &CreateRoomCmd{}, nil
}

func (c *CreateRoomCmd) HandleCommand(context *CmdContext) error {
	user := context.connection
	roomManager := context.roomManager

	if context.connection.Room != nil {
		return ErrUserAlreadyInRoom
	}

	userRoom := roomManager.CreateRoom()

	err := userRoom.Join(user.Token, user, user.UserName)
	if err != nil {
		roomManager.DeleteRoom(userRoom.Id)
		return err
	}

	user.Room = userRoom
	return nil
}
