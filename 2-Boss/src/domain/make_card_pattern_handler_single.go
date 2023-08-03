package domain

type MakeSingleCardPatternHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewMakeSingleCardPatternHandler(next *IMakeCardPatternHandler) *MakeSingleCardPatternHandler {
	return &MakeSingleCardPatternHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *MakeSingleCardPatternHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	ok = len(cards) == 1
	return NewSingleCardPattern(cards), ok
}
