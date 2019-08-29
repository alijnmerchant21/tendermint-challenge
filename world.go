package main

import (
	"fmt"
	"math/rand"
)

const MIN_STEPS_NUMBER = 10000

type World struct {
	Map
	aliens map[int]*Alien
	rand  *rand.Rand
}

func NewWorld(m *Map, n int, rand *rand.Rand) (*World, error) {
	nc := len(m.cities)
	if nc < n {
		fmt.Printf("given number of aliens (%v) is greater than number of cities (%v), only %v aliens will be used\n", n, nc, nc)
		n = nc
	}
	aliens := make(map[int]*Alien, n)
	cities := m.citiesAsArray()
	perm := rand.Perm(nc)
	for id := 1; id <= n; id++ { // starts from 1
		index := perm[id-1]
		c := cities[index]
		c.alienId = id
		aliens[id] = &Alien{id: id, cityName: c.name}
	}
	return &World{Map: *m, aliens: aliens, rand: rand}, nil
}

func (w *World) MoveAlien(id int) {
	fmt.Printf("alien#%d: start moving\n", id)
	cityName := w.aliens[id].cityName

	w.aliens[id].steps++

	city := w.cities[cityName]

	dirs := city.AvailableDirs()
	if len(dirs) == 0 {
		fmt.Printf("alien#%d: has no way to move\n", id)
		return
	}

	city.alienId = 0 // clean the city

	i := w.rand.Intn(len(dirs))
	dir := dirs[i]
	fmt.Printf("alien#%d: I can move to %v, moving to %v\n", id, dirs, dir)
	switch dir {
	case North:
		w.assingAlienToCity(id, city.north)
	case South:
		w.assingAlienToCity(id, city.south)
	case West:
		w.assingAlienToCity(id, city.west)
	case East:
		w.assingAlienToCity(id, city.east)
	}
}

func (w *World) assingAlienToCity(id int, cityName string) {
	w.aliens[id].cityName = cityName
	c := w.cities[cityName]
	if c.alienId != 0 {
		w.figth(id, c.alienId, c.name)
	} else {
		c.alienId = id
	}
}

func (w *World) figth(id1 int, id2 int, cityName string) {
	delete(w.aliens, id1)
	delete(w.aliens, id2)
	w.RemoveCity(cityName)
	fmt.Printf("> %s has been destroyed by alien %v and alien %v!\n", cityName, id1, id2)
}

func (w *World) StopCondition() bool {
	if len(w.aliens) == 0 {
		return true
	}
	for _, a := range w.aliens {
		if a.steps < MIN_STEPS_NUMBER {
			return false
		}
	}
	return true
}

func (w *World) String() string {
	return fmt.Sprintf("%v, %v", w.aliens, w.Map)
}
