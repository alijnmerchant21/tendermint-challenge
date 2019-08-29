package main

import (
	"bufio"
	"fmt"
	"io"
)

type Map struct {
	cities map[string]*City
}

func NewMap(r io.Reader) (*Map, error) {
	m := Map{cities: make(map[string]*City)}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		raw, err := ParseCity(line)
		if err != nil {
			return nil, err
		}
		putCityIntoMap(raw, m.cities)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &m, nil
}

func putCityIntoMap(raw *rawCity, cities map[string]*City) {
	city := City{name: raw.name}
	cities[raw.name] = &city
	if raw.north != "" {
		putIfNotExists(cities, raw.north)
		city.north = raw.north
		cities[raw.north].south = raw.name
	}
	if raw.south != "" {
		putIfNotExists(cities, raw.south)
		city.south = raw.south
		cities[raw.south].north = raw.name
	}
	if raw.west != "" {
		putIfNotExists(cities, raw.west)
		city.west = raw.west
		cities[raw.west].east = raw.name
	}
	if raw.east != "" {
		putIfNotExists(cities, raw.east)
		city.east = raw.east
		cities[raw.east].west = raw.name
	}
}

func putIfNotExists(cities map[string]*City, name string) {
	if _, ok := cities[name]; !ok {
		cities[name] = &City{name: name}
	}
}

func (m *Map) RemoveCity(name string) {
	c := m.cities[name]
	// fmt.Printf("%v\n", m.cities)
	if c.north != "" {
		m.cities[c.north].south = ""
	}
	if c.south != "" {
		m.cities[c.south].north = ""
	}
	if c.west != "" {
		m.cities[c.west].east = ""
	}
	if c.east != "" {
		m.cities[c.east].west = ""
	}
	delete(m.cities, name)
}

func (m Map) String() string {
	s := fmt.Sprintf("[")
	for _, c := range m.cities {
		s += fmt.Sprintf("%v,", c)
	}
	return s + "]"
}

func (m Map) citiesAsArray() [](*City) {
	arr := make([]*City, len(m.cities))
	i := 0
	for _, v := range m.cities {
		arr[i] = v
		i++
	}
	return arr
}
