package everodell

import (
	"encoding/csv"
	"os"
)

type Card struct {
	name         string
	unique       bool
	color        Color
	cost         Bundle
	partner      string
	points       int
	construction bool
	yield        Bundle

	occupied bool // if critter => partnered

	// action *Action
}

func (g *Game) readCards() {
	f, err := os.Open("data/cards.csv")

	defer f.Close()

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	g.deck = make([]Card, 0)

	for _, line := range data {
		num := readInt(line[6])

		for j := 0; j < num; j++ {
			g.deck = append(g.deck, Card{
				name:         line[0],
				unique:       readBool(line[1]),
				color:        readColor(line[2]),
				cost:         readBundle(line[3]),
				partner:      line[4],
				points:       readInt(line[5]),
				construction: readBool(line[7]),
				yield:        readBundle(line[8]),

				occupied: false,
			})
		}
	}
}

func (g *Game) canPlayCard(p *Player, c Card) int {
	if p.fullCity() {
		return 0 // 0 => can't play
	}

	for _, built := range p.city {
		if built.name == c.partner && !built.occupied {
			return 1 // => can play for free
		}
	}

	if p.resources.canPay(c.cost) {
		return 2 // => can play
	}

	return 0 // 0 => can't play
}

func (g *Game) playCard(p *Player, c Card) {
	n := g.canPlayCard(p, c)

	if n == 0 {
		return
	} else if n == 1 {
		for _, built := range p.city {
			if built.name == c.partner {
				built.occupied = true
				c.occupied = true
				break
			}
		}
	} else {
		p.resources.pay(c.cost)
	}

	p.city = append(p.city, c)

	g.processPlay(p, c)
}

func (g *Game) processPlay(p *Player, c Card) {
	switch c.color {
	case Tan:
	case Green:
		p.gain(c.yield) // non-standards will have yield "0T"
		g.trigger(p, c)
	case Blue:
	case Red:
	case Purple:
		return
	}
}

func (g *Game) trigger(p *Player, c Card) {
	switch c.name {

	}
}
