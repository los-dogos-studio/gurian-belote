package userconn

type StartGameCmd struct{}

func NewStartGameCmd(msg []byte) (Cmd, error) {
	return &StartGameCmd{}, nil
}

func (c *StartGameCmd) HandleCommand(context *CmdContext) error {
	user := context.connection

	if user.Room == nil {
		return ErrUserNotInRoom
	}

	return user.Room.StartGame()
}
