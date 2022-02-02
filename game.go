package everodell

type Game struct {
	players []*Player

	deck []Card
	meadow []Card

	events []Event
	actions []Action
}