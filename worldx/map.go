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

func putCityIntoMap(raw *City, cities map[string]*City) error {
	city := getOrCreate(cities, raw.name)
	for _, dir := range []Direction{North, South, West, East} {
		otherName := raw.getCityName(dir)
		if otherName != "" {
			other := getOrCreate(cities, otherName)
			if err := connectCities(city, dir, other); err != nil {
				return err
			}
		}
	}
	return nil
}

func getOrCreate(cities map[string]*City, name string) *City {
	if value, ok := cities[name]; ok {
		return value
	} else {
		c := &City{name: name}
		cities[name] = c
		return c
	}
}

func connectCities(src *City, dir Direction, dst *City) error {
	existent := src.getCityName(dir)
	if existent != "" && existent != dst.name {
		return fmt.Errorf("conflict: %s.%s = [%s, %s]", src.name, dir, existent, dst.name)
	}
	src.setDest(dir, dst.name)
	reverseExistent := dst.getCityName(dir.Opposite())
	if reverseExistent != "" && reverseExistent != src.name {
		return fmt.Errorf("conflict: %s.%s = [%s, %s]", dst.name, dir, reverseExistent, src.name)
	}
	dst.setDest(dir.Opposite(), src.name)
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
