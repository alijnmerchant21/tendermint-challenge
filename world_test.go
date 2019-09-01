package main

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
		{
			inputMap: &Map{map[string]*City{}},
			inputNum: 0,
			expectedCities: map[string]*City{},
			expectedAliens: map[int]*Alien{},
			hasError: false,
		},
		{
			inputMap: &Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}},
			},
			inputNum: 3, // given number is greater than number of cities
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
		{
			inputMap: &Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}},
			},
			inputNum: 1, // given number is less than number of cities
			expectedCities: map[string]*City{
				"Foo": &City{name: "Foo", alienId: 0 /* nobody here */, north: "Bar"},
				"Bar": &City{name: "Bar", alienId: 1, south: "Foo"},
			},
			expectedAliens: map[int]*Alien{
				1: &Alien{id: 1, cityName: "Bar"},
			},
			hasError: false,
		},
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
			inputMap: &Map{map[string]*City{}},
			inputNum: -1,
			expectedCities: map[string]*City{},
			expectedAliens: map[int]*Alien{},
			hasError: true,
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
