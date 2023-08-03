package domain

type ICollision struct {
	Collision
}

func NewICollision(c Collision) *ICollision {
	return &ICollision{
		Collision: c,
	}
}

func (ac *ICollision) handle(x1 int, x2 int) {
	if ac.match(x1, x2) {
		ac.doHandling(x1, x2)
	} else if ac.getNext() != nil {
		ac.getNext().handle(x1, x2)
	}
}

type Collision interface {
	match(x1 int, x2 int) bool
	doHandling(x1 int, x2 int)
	getNext() *ICollision
	setWorld(*World)
	getWorld() *World
}

type AbstractCollision struct {
	next  *ICollision
	world *World
}

func NewAbstractCollision(n *ICollision) *AbstractCollision {
	return &AbstractCollision{
		next: n,
	}
}

func (ac *AbstractCollision) getNext() *ICollision {
	return ac.next
}

func (ac *AbstractCollision) setWorld(w *World) {
	ac.world = w
	if ac.next != nil {
		ac.next.setWorld(w)
	}
}

func (ac *AbstractCollision) getWorld() *World {
	return ac.world
}
