package everodell

import (
	"encoding/csv"
	"os"
)

type Action struct {
	gains Bundle

	nUses    int
	nWorkers int

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
			gains:    readBundle(line[0]),
			nUses:    readInt(line[2]),
			nWorkers: 0,
		}

		if line[1] == "Yes" {
			specialActions[i] = action
		} else {
			g.actions[i] = action
		}
	}

	// choose Forest locations
	g.actions = append(g.actions, sample(4, specialActions)...)

	// todo journey, haven, etc
}

func (g *Game) visitAction(p *Player, a *Action) {
	if a.special != nil {
		a.special() // todo
	} else {
		p.gain(a.gains)
		p.nWorkers--
		a.nWorkers++
	}
}

func (g *Game) canVisitAction(p *Player, a *Action) bool {
	if p.nWorkers < 1 { // has no workers
		return false
	}

	switch a.nUses {
	case 0: // 0 => open action
		return true
	case 1: // 1 => closed action
		return a.nWorkers < 1
	case 4: // 4 => forest action, second is open for 4 players
		if len(g.players) < 4 {
			return a.nWorkers < 1
		}
		return a.nWorkers < 2
	}

	panic(a.nUses)
	return false // no action should get here
}
