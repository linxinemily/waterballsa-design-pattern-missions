package domain

type TankMoveBackwardCommand struct {
	tank *Tank
}

func NewTankMoveBackwardCommand(tank *Tank) *TankMoveBackwardCommand {
	return &TankMoveBackwardCommand{tank: tank}
}

func (cmd *TankMoveBackwardCommand) executeWrappedCmd() {
	cmd.tank.moveBackward()
}

func (cmd *TankMoveBackwardCommand) undoWrappedCmd() {
	cmd.tank.moveForward()
}

func (cmd *TankMoveBackwardCommand) Name() string {
	return "MoveTankBackward"
}
