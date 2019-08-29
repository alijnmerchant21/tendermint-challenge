package main

import (
	"bytes"
	"testing"
	"reflect"
	// "github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	expected Map
	hasError bool
}

func TestNewMap(t *testing.T) {
	var b []byte
	var m *Map
	var err error
	// var expected map[string]*City

	testCases := []testCase{
		{ ``, Map{map[string]*City{}}, false },
		// {City{name: "Foo", north: some}, []Direction{North}},
		// {City{name: "Foo", south: some}, []Direction{South}},
		// {City{name: "Foo", west: some}, []Direction{West}},
		// {City{name: "Foo", east: some}, []Direction{East}},
		// {City{name: "Foo", north: some, south: some, west: some, east: some}, []Direction{North, South, West, East}},
	}
	for _, tc := range testCases {
		b = []byte(tc.input)
		m, err = NewMap(bytes.NewReader(b))
		if tc.hasError {
			if err == nil {
				t.Errorf("NewMap('%s'): FAILED, expected an error but got the value '%v'", tc.input, m)
			} else {
				t.Logf("NewMap('%s'): PASSED, expected an error and got an error '%v'", tc.input, err)
			}
		} else {
			if len(tc.expected.cities) != len(m.cities) {
				t.Errorf("NewMap('%s'): FAILED, expected %v, but got '%v'", tc.input, tc.expected, m)
			}
			for k,v := range m.cities {
				if tc.expected.cities[k].name != v.name {
					t.Errorf("NewMap('%s'): FAILED, expected %v, but got '%v'", tc.input, tc.expected, m)
				}
				dirs := []string{"North","South","West","East"}
				for _, dir := range dirs {
					reflect.ValueOf(&t).MethodByName(dir).Call([]reflect.Value{})
				}
				ncity := tc.expected.cities[k].north
				if ncity == nil {
					if v.north != nil {
						// error
					}
				} else {
					if v.north == nil {
						// error
					}
					if v.north.name != ncity.name {
						// error
					}
					if v.north.south == nil || v.north.south != v.north {
						// error
					}
				}


					
			}
		
			t.Logf("NewMap('%s'): PASSED", tc.input)
		}
	}

	// b = []byte("")
	// m, err = NewMap(bytes.NewReader(b))
	// assert.Nil(t, err)
	// assert.True(t, len(m.cities) == 0, "")
	
	// b = []byte("Foo north=Bee")
	// m, err = NewMap(bytes.NewReader(b))
	// assert.Nil(t, err)
	// foo, ok := m.cities["Foo"]
	// assert.True(t, ok, "There is no Foo")
	// bee, ok := m.cities["Bee"]
	// assert.True(t, ok, "There is no Bee")
	// assert.True(t, foo.name == "Foo", "Wrong Foo name")
	// assert.True(t, bee.name == "Bee", "Wrong Bee name")
	// assert.True(t, foo.north == bee, "Foo.north must lead to Bee")
	// assert.True(t, foo.south == nil && foo.west == nil && foo.east == nil &&
	// 	bee.north == nil && bee.west == nil && bee.east == nil, "Other dirs must be nil")
	
	// b = []byte("Foo north=Foo")
	// m, err = NewMap(bytes.NewReader(b))
	// assert.Nil(t, m)
	// if err == nil {
	// 	t.Errorf("NewMap('Foo north=Foo'): FAILED, expected an error but got the value '%v'", m)
	// } else {
	// 	t.Logf("NewMap('Foo north=Foo'): PASSED, expected an error and got an error '%v'", err)
	// }
	// assert.True(t, err != nil, "an error was expected, but it was not")
}

func checkCities(t *testing.T, tc testCase, result map[string]*City) {
	if len(tc.expected) != len(result) {
		t.Errorf("NewMap('%s'): FAILED, expected %v, but got '%v'", tc.input, tc.expected, result)
	}

	t.Logf("NewMap('%s'): PASSED", tc.input)
}
