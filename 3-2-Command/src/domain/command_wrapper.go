package domain

type CommandWrapper struct {
	next *CommandWrapper
	Command
}

func NewICommandWrapper(command Command) *CommandWrapper {
	return &CommandWrapper{
		Command: command,
	}
}

func (command *CommandWrapper) SetNext(next *CommandWrapper) {
	command.next = next
}

func (command *CommandWrapper) execute() {
	command.executeWrappedCmd()
	if command.next != nil {
		command.next.execute()
	}
}

func (command *CommandWrapper) undo() {
	command.undoWrappedCmd()
	if command.next != nil {
		command.next.undo()
	}
}
