package everodell

import (
	"encoding/csv"
	"os"
)

type Game struct {
	players []*Player

	deck   []Card
	meadow []Card

	events  []Event
	actions []Action
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
