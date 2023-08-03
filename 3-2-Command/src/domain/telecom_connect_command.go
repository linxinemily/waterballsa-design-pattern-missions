package domain

type TelecomConnectCommand struct {
	telecom *Telecom
}

func NewTelecomConnectCommand(telecom *Telecom) *TelecomConnectCommand {
	return &TelecomConnectCommand{telecom: telecom}
}

func (cmd *TelecomConnectCommand) executeWrappedCmd() {
	cmd.telecom.connect()
}

func (cmd *TelecomConnectCommand) undoWrappedCmd() {
	cmd.telecom.disconnect()
}

func (cmd *TelecomConnectCommand) Name() string {
	return "ConnectTelecom"
}
