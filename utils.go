package everodell

type Season int
const (
	Winter Season = iota
	Spring
	Summer
	Autumn
)

type Component int
const (
	Wood Season = iota
	Resin
	Pebble
	Berry
)