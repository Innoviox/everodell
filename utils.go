package everodell

import (
	"math/rand"
	"strings"
)
import "strconv"

type Season int

const (
	Winter Season = iota
	Spring
	Summer
	Autumn
)

type Component int

const (
	Twig Component = iota
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
	nTwig   int
	nResin  int
	nPebble int
	nBerry  int
	nCards  int
	nPoints int
}

func (b *Bundle) add(o Bundle) {
	b.nTwig += o.nTwig
	b.nResin += o.nResin
	b.nPebble += o.nPebble
	b.nBerry += o.nBerry
	b.nCards += o.nCards // todo max 8 cards
	b.nPoints += o.nPoints
}

func readBundle(s string) Bundle {
	var b Bundle

	for _, s := range strings.Split(s, "/") {
		n, err := strconv.Atoi(string(s[0]))

		if err != nil {
			n = 1
		}

		switch s[len(s)-1:] {
		case "T":
			b.nTwig = n
		case "R":
			b.nResin = n
		case "P":
			b.nPebble = n
		case "B":
			b.nBerry = n
		case "C":
			b.nCards = n
		case "O":
			b.nPoints = n
		}
	}

	return b
}

func readInt(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		n = 0
	}

	return n
}

func sample[T](k int, arr []T) []T {
	var res []T

	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}

	return res
}
