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

		if err := putCityIntoMap(raw, m.cities); err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &m, nil
}

func putCityIntoMap(raw *rawCity, cities map[string]*City) error {
	city := City{name: raw.name}
	cities[raw.name] = &city
	if raw.north != "" {
		putIfNotExists(cities, raw.north)
		city.north = cities[raw.north]
		cities[raw.north].south = &city
	}
	if raw.south != "" {
		putIfNotExists(cities, raw.south)
		city.south = cities[raw.south]
		cities[raw.south].north = &city
	}
	if raw.west != "" {
		putIfNotExists(cities, raw.west)
		city.west = cities[raw.west]
		cities[raw.west].east = &city
	}
	if raw.east != "" {
		putIfNotExists(cities, raw.east)
		city.east = cities[raw.east]
		cities[raw.east].west = &city
	}
	return nil
}

func putIfNotExists(cities map[string]*City, name string) {
	if _, ok := cities[name]; !ok {
		cities[name] = &City{name: name}
	}
}

func (m *Map) RemoveCity(name string) {
	c := m.cities[name]
	// fmt.Printf("%v\n", m.cities)
	if c.north != nil {
		c.north.south = nil
	}
	if c.south != nil {
		c.south.north = nil
	}
	if c.west != nil {
		c.west.east = nil
	}
	if c.east != nil {
		c.east.west = nil
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
