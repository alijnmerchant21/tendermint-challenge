package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseCity(line string) (*City, error) {
	arr := strings.Fields(line)
	if len(arr) < 2 || len(arr) > 5 {
		return nil, fmt.Errorf("wrong line '%s'", line)
	}
	name, err := ParseCityName(arr[0])
	if err != nil {
		return nil, err
	}
	c := City{name: name}
	for _, v := range arr[1:] {
		dir := strings.Split(v, "=")
		if len(dir) != 2 {
			return nil, fmt.Errorf("wrong line '%s': wrong direction '%s'", line, dir)
		}
		if dir[1] == name {
			return nil, fmt.Errorf("wrong line '%s', wrong direction '%v': the road from the city '%s' can't lead to itself", line, dir, name)
		}
		switch dir[0] {
		case "north":
			c.north = dir[1]
		case "south":
			c.south = dir[1]
		case "west":
			c.west = dir[1]
		case "east":
			c.east = dir[1]
		default:
			return nil, fmt.Errorf("wrong line '%s': wrong direction '%s'", line, dir)
		}
	}
	return &c, nil
}

var forbidden = regexp.MustCompile(`[\d\W]+`)

func ParseCityName(s string) (string, error) {
	if len(s) == 0 || forbidden.MatchString(s) {
		return "", fmt.Errorf("wrong name '%s'", s)
	} else {
		return s, nil
	}
}
