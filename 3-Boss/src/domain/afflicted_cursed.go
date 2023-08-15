package domain

type Cursed struct {
	giver Role
	taker Role
}

func NewCursed(giver Role, taker Role) *Cursed {
	return &Cursed{
		giver: giver,
		taker: taker,
	}
}

func (c *Cursed) reward() {
	if c.giver.isAlive() {
		c.giver.addHp(c.taker.getMp())
	}
}
