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
		city, err := ParseCity(line)
		if err != nil {
			return nil, err
		}
		if err := putCityIntoMap(city, m.cities); err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &m, nil
}

func putCityIntoMap(new *City, cities map[string]*City) error {
	city := getOrCreate(cities, new.name)
	if new.north != "" {
		other := getOrCreate(cities, new.north)
		if city.north != "" && city.north != new.north {
			return fmt.Errorf("conflict: %s.north = [%s, %s]", city.name, city.north, new.north)
		}
		city.north = new.north
		if other.south != "" && other.south != new.name {
			return fmt.Errorf("conflict: %s.south = [%s, %s]", other.name, other.south, new.name)
		}
		other.south = new.name
	}
	if new.south != "" {
		other := getOrCreate(cities, new.south)
		if city.south != "" && city.south != new.south {
			return fmt.Errorf("conflict: %s.south = [%s, %s]", city.name, city.south, new.south)
		}
		city.south = new.south
		if other.north != "" && other.north != new.name {
			return fmt.Errorf("conflict: %s.north = [%s, %s]", other.name, other.north, new.name)
		}
		other.north = new.name
	}
	if new.west != "" {
		other := getOrCreate(cities, new.west)
		if city.west != "" && city.west != new.west {
			return fmt.Errorf("conflict: %s.west = [%s, %s]", city.name, city.west, new.west)
		}
		city.west = new.west
		if other.east != "" && other.east != new.name {
			return fmt.Errorf("conflict: %s.east = [%s, %s]", other.name, other.east, new.name)
		}
		other.east = new.name
	}
	if new.east != "" {
		other := getOrCreate(cities, new.east)
		if city.east != "" && city.east != new.east {
			return fmt.Errorf("conflict: %s.east = [%s, %s]", city.name, city.east, new.east)
		}
		city.east = new.east
		if other.west != "" && other.west != new.name {
			return fmt.Errorf("conflict: %s.west = [%s, %s]", other.name, other.west, new.name)
		}
		other.west = new.name
	}
	return nil
}

func getOrCreate(cities map[string]*City, name string) *City {
	if value, ok := cities[name]; ok {
		return value
	} else {
		new := &City{name: name}
		cities[name] = new
		return new
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
