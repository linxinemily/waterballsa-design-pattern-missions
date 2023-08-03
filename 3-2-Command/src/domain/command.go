package domain

type Command interface {
	executeWrappedCmd()
	undoWrappedCmd()
	Name() string
}
