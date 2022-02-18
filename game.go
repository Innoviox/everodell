package everodell

type Game struct {
	players []*Player

	deck   []Card
	meadow []Card

	events  []Event
	actions []Action
}

func NewGame() *Game {
	g := &Game{}

	g.deck = make([]Card, 0)
	g.readCards()

	g.players = make([]*Player, 0)

	g.meadow = g.DrawCards(8)

	g.events = make([]Event, 0)

	g.actions = make([]Action, 0)
	g.ReadActions()

	return g
}
