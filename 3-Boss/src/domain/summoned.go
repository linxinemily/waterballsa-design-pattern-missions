package domain

type Summoned struct {
	summoner Role
	summoned *Slime
}

func NewSummoned(summoner Role, summoned *Slime) *Summoned {
	return &Summoned{
		summoner: summoner,
		summoned: summoned,
	}
}

func (s *Summoned) reward() {
	if s.summoner.isAlive() {
		s.summoner.addHp(30)
	}
}
