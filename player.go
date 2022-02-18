package everodell

type Player struct {
	hand []Card

	city   []Card
	ghosts []Card // for wife & wanderer

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

func (g *Game) MakePlayer() *Player {
	return &Player{
		hand:   make([]Card, 8),
		city:   make([]Card, 15),
		ghosts: make([]Card, 0),
		resources: Bundle{
			nTwig:   0,
			nResin:  0,
			nPebble: 0,
			nBerry:  0,
			nCards:  0,
			nPoints: 0,
		},
		season:   Winter,
		nWorkers: 2,
	}
}

func (g *Game) NextSeason(p *Player) {
	p.season += 1
	switch p.season {
	case Spring:
		p.nWorkers += 1
		g.triggerGreens(p)
	case Summer:
		p.nWorkers += 1
		// todo draw 2 meadow cards
	case Autumn:
		p.nWorkers += 2
		g.triggerGreens(p)
	}
}
