package userconn

type CreateRoomCmd struct {
}

func NewCreateRoomCmd(msg []byte) (Cmd, error) {
	return &CreateRoomCmd{}, nil
}

func (c *CreateRoomCmd) HandleCommand(context *CmdContext) error {
	user := context.user
	roomManager := context.roomManager

	if context.user.Room != nil {
		return ErrUserAlreadyInRoom
	}

	userRoom := roomManager.CreateRoom()

	err := userRoom.Join(user.UserId, user)
	if err != nil {
		roomManager.DeleteRoom(userRoom.Id)
		return err
	}

	user.Room = userRoom
	return nil
}
