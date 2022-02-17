package everodell

type Player struct {
	hand []Card

	resources Bundle
	season    Season
	nWorkers  int
}

func (p *Player) gain(b Bundle) {
	p.resources.add(b)
	// todo draw cards
}
