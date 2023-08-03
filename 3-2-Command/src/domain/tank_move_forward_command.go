package domain

type TankMoveForwardCommand struct {
	tank *Tank
}

func NewTankMoveForwardCommand(tank *Tank) *TankMoveForwardCommand {
	return &TankMoveForwardCommand{tank: tank}
}

func (cmd *TankMoveForwardCommand) executeWrappedCmd() {
	cmd.tank.moveForward()
}

func (cmd *TankMoveForwardCommand) undoWrappedCmd() {
	cmd.tank.moveBackward()
}

func (cmd *TankMoveForwardCommand) Name() string {
	return "MoveTankForward"
}
