package main

import (
	"fmt"
)

type City struct {
	name string

	alienId int // 0 means nobody!

	north string // empty means there is no way
	south string
	west  string
	east  string
}

func (c *City) AvailableDirs() []Direction {
	dirs := make([]Direction, 0)
	if c.north != "" {
		dirs = append(dirs, North)
	}
	if c.south != "" {
		dirs = append(dirs, South)
	}
	if c.west != "" {
		dirs = append(dirs, West)
	}
	if c.east != "" {
		dirs = append(dirs, East)
	}
	return dirs
}

func (c City) String() string {
	s := c.name
	if c.north != "" {
		s += fmt.Sprintf(" north=%s", c.north)
	}
	if c.south != "" {
		s += fmt.Sprintf(" south=%s", c.south)
	}
	if c.west != "" {
		s += fmt.Sprintf(" west=%s", c.west)
	}
	if c.east != "" {
		s += fmt.Sprintf(" east=%s", c.east)
	}
	return s
}
