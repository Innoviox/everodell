package everodell

type Card struct {
	name string
	cost []Component

	points int

	action *Action
}