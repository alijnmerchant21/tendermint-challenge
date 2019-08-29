package main

import (
	"fmt"
)

type City struct {
	name string

	alienId int // 0 means nobody!

	north *City
	south *City
	west  *City
	east  *City
}

func (c *City) AvailableDirs() []Direction {
	dirs := make([]Direction, 0)
	if c.north != nil {
		dirs = append(dirs, North)
	}
	if c.south != nil {
		dirs = append(dirs, South)
	}
	if c.west != nil {
		dirs = append(dirs, West)
	}
	if c.east != nil {
		dirs = append(dirs, East)
	}
	return dirs
}

func (c City) String() string {
	s := fmt.Sprintf("{%s (%v)", c.name, c.alienId)
	if c.north != nil {
		s += fmt.Sprintf(" north=%s", c.north.name)
	}
	if c.south != nil {
		s += fmt.Sprintf(" south=%s", c.south.name)
	}
	if c.west != nil {
		s += fmt.Sprintf(" west=%s", c.west.name)
	}
	if c.east != nil {
		s += fmt.Sprintf(" east=%s", c.east.name)
	}
	return s + "}"
}
