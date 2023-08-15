package domain

type Summoned struct {
	summoner Role
	summoned Role
}

func NewSummoned(summoner Role, summoned Role) *Summoned {
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
