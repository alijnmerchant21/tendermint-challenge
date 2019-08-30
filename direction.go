package main

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	West  Direction = "west"
	East  Direction = "east"
)

var pairs = map[Direction]Direction{
	North: South,
	South: North,
	West:  East,
	East:  West,
}

func (d Direction) Opposite() Direction {
	return pairs[d]
}
