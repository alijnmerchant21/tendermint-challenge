package worldx

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
	city := &City{name: name}
	for _, v := range arr[1:] {
		dir := strings.Split(v, "=")
		if len(dir) != 2 {
			return nil, fmt.Errorf("wrong line '%s': wrong direction '%s'", line, dir)
		}
		other, err := ParseCityName(dir[1])
		if err != nil {
			return nil, err
		}
		if other == name {
			return nil, fmt.Errorf("wrong line '%s', wrong direction '%v': the road from the city '%s' can't lead to itself", line, dir, name)
		}
		switch dir[0] {
		case "north":
			city.north = other
		case "south":
			city.south = other
		case "west":
			city.west = other
		case "east":
			city.east = other
		default:
			return nil, fmt.Errorf("wrong line '%s': wrong direction '%s'", line, dir)
		}
	}
	return city, nil
}

var allowed = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-]*[a-zA-Z0-9]$`)

func ParseCityName(s string) (string, error) {
	if allowed.MatchString(s) {
		return s, nil
	} else {
		return "", fmt.Errorf("wrong name '%s'", s)
	}
}
