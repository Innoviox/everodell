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

	occupied bool

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

	for i, line := range data {
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

				occupied: false,
			})
		}
	}
}
