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
	Wood Component = iota
	Resin
	Pebble
	Berry
)

type Color int
const (
	Tan Color = iota
	Green
	Red
	Blue
	Purple
)

type Bundle struct {
	nWood int
	nResin int
	nPebble int
	nBerry int
}