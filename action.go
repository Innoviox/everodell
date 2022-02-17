package everodell

import (
	"encoding/csv"
	"os"
)

type Action struct {
	gains Bundle

	nUses int

	special func() // todo
}

func (g *Game) ReadActions() {
	f, err := os.Open("data/actions.csv")

	defer f.Close()

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	specialActions := make([]Action, len(data))
	g.actions = make([]Action, len(data))

	for i, line := range data {
		if i == 0 { // header
			continue
		}

		action := Action{
			gains: readBundle(line[0]),
			nUses: readInt(line[2]),
		}

		if line[1] == "Yes" {
			specialActions[i] = action
		} else {
			g.actions[i] = action
		}
	}

	// choose Forest locations
	g.actions = append(g.actions, sample(4, specialActions)...)
}

func (g *Game) visitActions(p *Player, a Action) {
	if a.special != nil {
		a.special()
	} else {
		p.gain(a.gains)
	}
}
