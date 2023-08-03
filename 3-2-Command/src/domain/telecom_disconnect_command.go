package domain

type TelecomDisconnectCommand struct {
	telecom *Telecom
}

func NewTelecomDisconnectCommand(telecom *Telecom) *TelecomDisconnectCommand {
	return &TelecomDisconnectCommand{telecom: telecom}
}

func (cmd *TelecomDisconnectCommand) executeWrappedCmd() {
	cmd.telecom.disconnect()
}

func (cmd *TelecomDisconnectCommand) undoWrappedCmd() {
	cmd.telecom.connect()
}

func (cmd *TelecomDisconnectCommand) Name() string {
	return "DisconnectTelecom"
}
