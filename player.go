package everodell

type Player struct {
	hand []Card

	city []Card

	resources Bundle
	season    Season
	nWorkers  int
}

func (p *Player) gain(b Bundle) {
	p.resources.add(b)
	// todo draw cards
}

func (p *Player) fullCity() bool {
	return len(p.city) == 15
}

func (p *Player) fullHand() bool {
	return len(p.hand) == 8
}
