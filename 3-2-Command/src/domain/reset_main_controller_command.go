package domain

type ResetMainControllerKeyboardCommand struct {
	ctr    *MainController
	backup map[byte]*CommandWrapper
}

func NewResetMainControllerKeyboardCommand(ctr *MainController) *ResetMainControllerKeyboardCommand {
	return &ResetMainControllerKeyboardCommand{ctr: ctr}
}

func (cmd *ResetMainControllerKeyboardCommand) executeWrappedCmd() {
	cmd.backup = cmd.ctr.bindMap
	cmd.ctr.Reset()
}

func (cmd *ResetMainControllerKeyboardCommand) undoWrappedCmd() {
	cmd.ctr.bindMap = cmd.backup
}

func (cmd *ResetMainControllerKeyboardCommand) Name() string {
	return "ResetMainControlKeyboard"
}
