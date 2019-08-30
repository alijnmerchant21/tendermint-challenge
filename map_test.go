package main

import (
	"bytes"
	"testing"
)

func TestNewMap(t *testing.T) {
	testCases := []struct {
		input    string
		expected Map
		hasError bool
	}{
		// parsing errors:
		{"Foo", Map{}, true},
		{"Foo north=Foo", Map{}, true},
		{"Foo wrongDirection=Bar", Map{}, true},

		// conflicting errors:
		{"Foo north=Bar\nBar south=Bee", Map{}, true},
		{"Foo north=Bar\nBaz south=Foo", Map{}, true},

		{"Foo south=Bar\nBar north=Bee", Map{}, true},
		{"Foo south=Bar\nBaz north=Foo", Map{}, true},

		{"Foo west=Bar\nBar east=Bee", Map{}, true},
		{"Foo west=Bar\nBaz east=Foo", Map{}, true},

		{"Foo east=Bar\nBar west=Bee", Map{}, true},
		{"Foo east=Bar\nBaz west=Foo", Map{}, true},

		// happy paths:
		{
			input:    "",
			expected: Map{map[string]*City{}},
			hasError: false},
		{
			input: "Foo north=Bar",
			expected: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}}},
			hasError: false},
		{
			input: "Foo north=Bar\nBar south=Foo",
			expected: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}}},
			hasError: false},
		{
			input: "Foo north=Bar west=Baz south=Qux\nBar south=Foo west=Bee",
			expected: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar", west: "Baz", south: "Qux"},
				"Bar": &City{name: "Bar", south: "Foo", west: "Bee"},
				"Baz": &City{name: "Baz", east: "Foo"},
				"Qux": &City{name: "Qux", north: "Foo"},
				"Bee": &City{name: "Bee", east: "Bar"}}},
			hasError: false},
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

func TestRemoveCity(t *testing.T) {
	testCases := []struct {
		receiver Map
		input    string
		expected Map
	}{
		{
			receiver: Map{},
			input:    "any",
			expected: Map{},
		},
		{
			receiver: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar"},
				"Bar": &City{name: "Bar", south: "Foo"}}},
			input: "Foo",
			expected: Map{map[string]*City{
				"Bar": &City{name: "Bar"}}},
		},
		{
			receiver: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar", east: "Bee"},
				"Bar": &City{name: "Bar", south: "Foo"},
				"Bee": &City{name: "Bee", west: "Foo"}}},
			input: "Foo",
			expected: Map{map[string]*City{
				"Bar": &City{name: "Bar"},
				"Bee": &City{name: "Bee"}}},
		},
		{
			receiver: Map{map[string]*City{
				"Foo": &City{name: "Foo", north: "Bar", east: "Bee"},
				"Bar": &City{name: "Bar", south: "Foo"},
				"Bee": &City{name: "Bee", west: "Foo"}}},
			input: "Bar",
			expected: Map{map[string]*City{
				"Foo": &City{name: "Foo", east: "Bee"},
				"Bee": &City{name: "Bee", west: "Foo"}}},
		},
	}
	for _, tc := range testCases {
		original := tc.receiver.String()
		tc.receiver.RemoveCity(tc.input)
		if len(tc.expected.cities) != len(tc.receiver.cities) {
			t.Errorf("%v.RemoveCity('%s'): FAILED, expected len=%v, but got len=%v", original, tc.input, len(tc.expected.cities), len(tc.receiver.cities))
		}
		for k, actualCity := range tc.receiver.cities {
			expectedCity := *tc.expected.cities[k]
			if expectedCity != *actualCity {
				t.Errorf("%v.RemoveCity('%s'): FAILED, expected %v, but got '%v'", original, tc.input, expectedCity, *actualCity)
			}
		}
		t.Logf("%v.RemoveCity('%s'): PASSED", original, tc.input)
	}

}
