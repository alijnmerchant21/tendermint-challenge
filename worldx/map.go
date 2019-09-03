package worldx

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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
		if err := connectCities(city, North, other); err != nil {
			return err
		}
	}
	if new.south != "" {
		other := getOrCreate(cities, new.south)
		if err := connectCities(city, South, other); err != nil {
			return err
		}
	}
	if new.west != "" {
		other := getOrCreate(cities, new.west)
		if err := connectCities(city, West, other); err != nil {
			return err
		}
	}
	if new.east != "" {
		other := getOrCreate(cities, new.east)
		if err := connectCities(city, East, other); err != nil {
			return err
		}
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

func connectCities(from *City, dir Direction, to *City) error {
	if from.dest(dir) != "" && from.dest(dir) != to.name {
		return fmt.Errorf("conflict: %s.%s = [%s, %s]", from.name, dir, from.dest(dir), to.name)
	}
	from.setDest(dir, to.name)
	if to.dest(dir.Opposite()) != "" && to.dest(dir.Opposite()) != from.name {
		return fmt.Errorf("conflict: %s.%s = [%s, %s]", to.name, dir, to.dest(dir.Opposite()), from.name)
	}
	to.setDest(dir.Opposite(), from.name)
	return nil
}

func (m *Map) RemoveCity(name string) {
	c, ok := m.cities[name]
	if !ok {
		return
	}
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

func (m Map) String() (s string) {
	for _, c := range m.cities {
		s += fmt.Sprintf("%v\n", c)
	}
	return
}

// returns sorted array to provide determinism in tests
// todo: use only one aray maybe...
func (m Map) citiesAsSortedArray() []*City {
	arr := make([]*City, len(m.cities))
	keys := m.sortedArrayOfCityNames()
	for i, k := range keys {
		arr[i] = m.cities[k]
	}
	return arr
}

func (m Map) sortedArrayOfCityNames() []string {
	keys := make([]string, len(m.cities))
	j := 0
	for k := range m.cities {
		keys[j] = k
		j++
	}
	sort.Strings(keys)
	return keys
}
