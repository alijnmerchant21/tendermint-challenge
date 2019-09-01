package worldx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorld(t *testing.T) {
	testCases := []struct {
		inputMap       *Map
		inputNum       int
		expectedCities map[string]*City
		expectedAliens map[int]*Alien
		hasError       bool
	}{
		// empty map
		{
			inputMap:       &Map{map[string]*City{}},
			inputNum:       0,
			expectedCities: map[string]*City{},
			expectedAliens: map[int]*Alien{},
			hasError:       false,
		},

		// given aliens number is greater than number of cities
		{
			inputMap: &Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}},
			},
			inputNum: 3,
			expectedCities: map[string]*City{
				"Foo": &City{name: "Foo", alienId: 2, north: "Bar"},
				"Bar": &City{name: "Bar", alienId: 1, south: "Foo"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "Bar"},
				2: &Alien{id: 2, cityName: "Foo"},
			},
			hasError: false,
		},

		// given aliens number is less than number of cities
		{
			inputMap: &Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}},
			},
			inputNum: 1,
			expectedCities: map[string]*City{
				"Foo": &City{name: "Foo", alienId: 0 /* nobody here */, north: "Bar"},
				"Bar": &City{name: "Bar", alienId: 1, south: "Foo"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "Bar"},
			},
			hasError: false,
		},

		// check at least somehow that aliens are distributed by map
		{
			inputMap: &Map{map[string]*City{
				"A": &City{name: "A"},
				"B": &City{name: "B"},
				"C": &City{name: "C"},
				"D": &City{name: "D"},
				"E": &City{name: "E"},
				"F": &City{name: "F"}},
			},
			inputNum: 6,
			expectedCities: map[string]*City{
				"A": &City{name: "A", alienId: 1},
				"B": &City{name: "B", alienId: 2},
				"C": &City{name: "C", alienId: 3},
				"D": &City{name: "D", alienId: 4},
				"E": &City{name: "E", alienId: 5},
				"F": &City{name: "F", alienId: 6},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "A"},
				2: &Alien{id: 2, cityName: "B"},
				3: &Alien{id: 3, cityName: "C"},
				4: &Alien{id: 4, cityName: "D"},
				5: &Alien{id: 5, cityName: "E"},
				6: &Alien{id: 6, cityName: "F"},
			},
			hasError: false,
		},
		{
			inputMap:       &Map{map[string]*City{}},
			inputNum:       -1,
			expectedCities: map[string]*City{},
			expectedAliens: map[int]*Alien{},
			hasError:       true,
		},
	}

	for _, tc := range testCases {
		r := FakeRandomizer{}
		w, err := NewWorld(tc.inputMap, tc.inputNum, r)
		if tc.hasError {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, len(w.aliens), len(tc.expectedAliens))
			assert.Equal(t, len(w.cities), len(tc.expectedCities))
			for k, actualCity := range w.cities {
				expectedCity := tc.expectedCities[k]
				assert.Equal(t, *expectedCity, *actualCity)
			}
			for k, actualAlien := range w.aliens {
				expectedAlien := tc.expectedAliens[k]
				assert.Equal(t, *expectedAlien, *actualAlien)
			}
		}
	}
}

func TestMoveAlien(t *testing.T) {
	testCases := []struct {
		world          World
		alienId        int
		expectedCities map[string]*City
		expectedAliens map[int]*Alien
		hasError       bool
	}{
		// empty world
		{
			world:          World{Map: Map{map[string]*City{}}, aliens: map[int]*Alien{}},
			alienId:        0,
			expectedCities: map[string]*City{},
			expectedAliens: map[int]*Alien{},
			hasError:       true,
		},

		// alien goes north
		{
			world:   createCrossWorld(FakeRandomizer{fakeIntn: 0}),
			alienId: 1,
			expectedCities: map[string]*City{
				"Center": &City{name: "Center", alienId: 0, north: "N", south: "S", west: "W", east: "E"},
				"N":      &City{name: "N", alienId: 1, south: "Center"},
				"S":      &City{name: "S", alienId: 0, south: "Center"},
				"W":      &City{name: "W", alienId: 0, south: "Center"},
				"E":      &City{name: "E", alienId: 0, south: "Center"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "N", steps: 1},
			},
			hasError: false,
		},

		// alien goes south
		{
			world:   createCrossWorld(FakeRandomizer{fakeIntn: 1}),
			alienId: 1,
			expectedCities: map[string]*City{
				"Center": &City{name: "Center", alienId: 0, north: "N", south: "S", west: "W", east: "E"},
				"N":      &City{name: "N", alienId: 0, south: "Center"},
				"S":      &City{name: "S", alienId: 1, south: "Center"},
				"W":      &City{name: "W", alienId: 0, south: "Center"},
				"E":      &City{name: "E", alienId: 0, south: "Center"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "S", steps: 1},
			},
			hasError: false,
		},

		// alien goes west
		{
			world:   createCrossWorld(FakeRandomizer{fakeIntn: 2}),
			alienId: 1,
			expectedCities: map[string]*City{
				"Center": &City{name: "Center", alienId: 0, north: "N", south: "S", west: "W", east: "E"},
				"N":      &City{name: "N", alienId: 0, south: "Center"},
				"S":      &City{name: "S", alienId: 0, south: "Center"},
				"W":      &City{name: "W", alienId: 1, south: "Center"},
				"E":      &City{name: "E", alienId: 0, south: "Center"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "W", steps: 1},
			},
			hasError: false,
		},

		// alien goes east
		{
			world:   createCrossWorld(FakeRandomizer{fakeIntn: 3}),
			alienId: 1,
			expectedCities: map[string]*City{
				"Center": &City{name: "Center", alienId: 0, north: "N", south: "S", west: "W", east: "E"},
				"N":      &City{name: "N", alienId: 0, south: "Center"},
				"S":      &City{name: "S", alienId: 0, south: "Center"},
				"W":      &City{name: "W", alienId: 0, south: "Center"},
				"E":      &City{name: "E", alienId: 1, south: "Center"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "E", steps: 1},
			},
			hasError: false,
		},

		// alien has nowhere to go
		{
			world: World{
				Map: Map{map[string]*City{
					"A": &City{name: "A", alienId: 1},
					"B": &City{name: "B", alienId: 0, north: "C"},
					"C": &City{name: "C", alienId: 0, south: "B"},
				}},
				aliens: map[int]*Alien{
					1: &Alien{id: 1, cityName: "A", steps: 0},
				},
				rand: FakeRandomizer{},
			},
			alienId: 1,
			expectedCities: map[string]*City{
				"A": &City{name: "A", alienId: 1},
				"B": &City{name: "B", alienId: 0, north: "C"},
				"C": &City{name: "C", alienId: 0, south: "B"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "A", steps: 1},
			},
			hasError: false,
		},

		// alien goes to a city that is occupied by another
		{
			world: World{
				Map: Map{map[string]*City{
					"A": &City{name: "A", alienId: 1, south: "B"},
					"B": &City{name: "B", alienId: 2, north: "A"},
				}},
				aliens: map[int]*Alien{
					1: &Alien{id: 1, cityName: "A", steps: 0},
					2: &Alien{id: 2, cityName: "B", steps: 0},
				},
				rand: FakeRandomizer{},
			},
			alienId: 1,
			expectedCities: map[string]*City{
				"A": &City{name: "A", alienId: 0},
			},
			expectedAliens: map[int]*Alien{},
			hasError:       false,
		},
	}

	for _, tc := range testCases {
		err := tc.world.MoveAlien(tc.alienId)
		if tc.hasError {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, len(tc.world.aliens), len(tc.expectedAliens))
			assert.Equal(t, len(tc.world.cities), len(tc.expectedCities))
			for k, actualCity := range tc.world.cities {
				expectedCity := tc.expectedCities[k]
				assert.Equal(t, *expectedCity, *actualCity)
			}
			for k, actualAlien := range tc.world.aliens {
				expectedAlien := tc.expectedAliens[k]
				assert.Equal(t, *expectedAlien, *actualAlien)
			}
		}
	}
}

func createCrossWorld(rand Random) World {
	return World{
		Map: Map{map[string]*City{
			"Center": &City{name: "Center", alienId: 1, north: "N", south: "S", west: "W", east: "E"},
			"N":      &City{name: "N", alienId: 0, south: "Center"},
			"S":      &City{name: "S", alienId: 0, south: "Center"},
			"W":      &City{name: "W", alienId: 0, south: "Center"},
			"E":      &City{name: "E", alienId: 0, south: "Center"},
		}},
		aliens: map[int]*Alien{
			1: &Alien{id: 1, cityName: "Center", steps: 0},
		},
		rand: rand,
	}
}

func TestStopCondition(t *testing.T) {
	testCases := []struct {
		world          World
		expectedResult bool
	}{
		// no aliens
		{
			world: World{
				Map:    Map{map[string]*City{}},
				aliens: map[int]*Alien{},
				rand:   FakeRandomizer{},
			},
			expectedResult: true,
		},

		// one alien who's just started
		{
			world: World{
				Map: Map{map[string]*City{
					"A": &City{name: "A", alienId: 1},
				}},
				aliens: map[int]*Alien{
					1: &Alien{id: 1, cityName: "A", steps: 0},
				},
				rand: FakeRandomizer{},
			},
			expectedResult: false,
		},

		// one alien who just walked 9999 steps, and one who walked 10K steps
		{
			world: World{
				Map: Map{map[string]*City{
					"A": &City{name: "A", alienId: 1},
					"B": &City{name: "B", alienId: 2},
				}},
				aliens: map[int]*Alien{
					1: &Alien{id: 1, cityName: "A", steps: 9999},
					2: &Alien{id: 2, cityName: "B", steps: 10000},
				},
				rand: FakeRandomizer{},
			},
			expectedResult: false,
		},

		// two aliens who have completed 10K steps
		{
			world: World{
				Map: Map{map[string]*City{
					"A": &City{name: "A", alienId: 1},
					"B": &City{name: "B", alienId: 2},
				}},
				aliens: map[int]*Alien{
					1: &Alien{id: 1, cityName: "A", steps: 10000},
					2: &Alien{id: 2, cityName: "B", steps: 10000},
				},
				rand: FakeRandomizer{},
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		res := tc.world.stopCondition()
		assert.Equal(t, tc.expectedResult, res)
	}
}
