package everodell

type Card struct {
	name string
	cost Bundle

	points int

	color Color
	unique bool
	construction bool
	partner string

	// action *Action
}