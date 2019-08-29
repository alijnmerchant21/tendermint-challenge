package main

import (
	"bytes"
	"testing"
)

type testCase struct {
	input string
	expected Map
	hasError bool
}

func TestNewMap(t *testing.T) {
	testCases := []testCase{
		// errors:
		{ "Foo", Map{}, true },
		{ "Foo north=Foo", Map{}, true },
		{ "Foo wrongDirection=Bar", Map{}, true },

		// happy paths:
		{ "",
			Map{map[string]*City{}}, false },
		{ "Foo north=Bar",
			Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}}}, false },
		{ "Foo north=Bar\nBar south=Foo",
			Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}}}, false },
		{ "Foo north=Bar west=Baz south=Qux\nBar south=Foo west=Bee",
			Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar", west: "Baz", south: "Qux"},
				"Bar": &City{name: "Bar", south: "Foo", west: "Bee"},
				"Baz": &City{name: "Baz", east: "Foo"},
				"Qux": &City{name: "Qux", north: "Foo"},
				"Bee": &City{name: "Bee", east: "Bar"}}}, false },
	}
	for _, tc := range testCases {
		b := []byte(tc.input)
		m, err := NewMap(bytes.NewReader(b))
		if tc.hasError {
			if err == nil {
				t.Errorf("NewMap from '%s': FAILED, expected an error but got the value '%v'", tc.input, m)
			} else {
				t.Logf("NewMap from '%s': PASSED, expected an error and got an error '%v'", tc.input, err)
			}
		} else {
			if len(tc.expected.cities) != len(m.cities) {
				t.Errorf("NewMap from '%s': FAILED, expected len=%v, but got '%v', len=%v", tc.input, len(tc.expected.cities), m, len(m.cities))
			}
			for k, actualCity := range m.cities {
				expectedCity := *tc.expected.cities[k]
				if expectedCity != *actualCity {
					t.Errorf("NewMap from '%s': FAILED, expected %v, but got '%v'", tc.input, expectedCity, *actualCity)
				}
			}
			t.Logf("NewMap from '%s': PASSED", tc.input)
		}
	}

}
