package domain

type IMakeCardPatternHandler struct {
	MakeCardPatternHandler
}

func NewIMakeCardPatternHandler(handler MakeCardPatternHandler) *IMakeCardPatternHandler {
	return &IMakeCardPatternHandler{
		MakeCardPatternHandler: handler,
	}
}

func (handler *IMakeCardPatternHandler) handle(cards []*Card) CardPattern { // template method

	if cardPattern, ok := handler.match(cards); ok {
		return cardPattern
	} else if handler.getNext() != nil {
		return handler.getNext().handle(cards)
	}

	return nil
}

type MakeCardPatternHandler interface {
	match(cards []*Card) (cardPattern CardPattern, ok bool)
	getNext() *IMakeCardPatternHandler
}

type AbstractMakeCardPatternHandler struct {
	cardPatternType string
	next            *IMakeCardPatternHandler
}

func NewAbstractHandler(next *IMakeCardPatternHandler) *AbstractMakeCardPatternHandler {
	return &AbstractMakeCardPatternHandler{
		next: next,
	}
}

func (handler *AbstractMakeCardPatternHandler) getNext() *IMakeCardPatternHandler {
	return handler.next
}
