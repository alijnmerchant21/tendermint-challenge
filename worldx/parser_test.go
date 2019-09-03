package worldx

import (
	"testing"
)

func TestParseCity(t *testing.T) {
	type TestDataItem struct {
		input    string
		result   *City
		hasError bool
	}
	testDataItems := []TestDataItem{
		{"", nil, true},
		{"1", nil, true},
		{"Foo", nil, true},
		{"Foo Bee", nil, true},
		{"Foo nor=Bee", nil, true},
		{"wrong/name north=Bee", nil, true},
		{"Foo north=wrong/name", nil, true},
		{"Foo south=wrong/name", nil, true},
		{"Foo west=wrong/name", nil, true},
		{"Foo east=wrong/name", nil, true},
		{"Foo north=Bee", &City{name: "Foo", north: "Bee"}, false},
		{"Foo north=Bee south=Bar", &City{name: "Foo", north: "Bee", south: "Bar"}, false},
		{"Foo north=Bee south=Bar west=Baz", &City{name: "Foo", north: "Bee", south: "Bar", west: "Baz"}, false},
		{"Foo north=Bee south=Bar west=Baz east=Cat", &City{name: "Foo", north: "Bee", south: "Bar", west: "Baz", east: "Cat"}, false},
		{"Foo west=Baz east=Cat north=Bee south=Bar", &City{name: "Foo", north: "Bee", south: "Bar", west: "Baz", east: "Cat"}, false},
		{"Foo west=Baz east=Cat north=Bee south=Bar north=Bee", nil, true},

		{"Foo south=Foo", nil, true}, // we can't go from Foo to Foo
	}
	for _, item := range testDataItems {
		result, err := ParseCity(item.input)
		if item.hasError {
			// expected an error
			if err == nil {
				t.Errorf("ParseCity('%v'): FAILED, expected an error but got the value '%v'", item.input, result)
			} else {
				t.Logf("ParseCity('%v'): PASSED, expected an error and got an error '%v'", item.input, err)
			}
		} else {
			// expected a value
			if *result != *item.result {
				t.Errorf("ParseCity('%v'): FAILED, expected '%v' but got value '%v'", item.input, *item.result, *result)
			} else {
				t.Logf("ParseCity('%v'): PASSED", item.input)
			}
		}
	}
}

func TestParseCityName(t *testing.T) {
	type TestDataItem struct {
		input    string
		result   string
		hasError bool
	}
	testDataItems := []TestDataItem{
		{"", "", true},
		{"-1", "", true},
		{"123", "", true},
		{"bar.", "", true},
		{"1bar", "", true},
		{"-bar", "", true},
		{"bar foo", "", true},
		{"bar-", "bar-", true},
		{"bar1", "bar1", false},
		{"bar1bee", "bar1bee", false},
		{"bar-foo", "bar-foo", false},
		{"bar", "bar", false},
		{"Bar", "Bar", false},
		{"aa", "aa", false},
		{"a", "a", false},
	}
	for _, item := range testDataItems {
		result, err := ParseCityName(item.input)
		if item.hasError {
			// expected an error
			if err == nil {
				t.Errorf("ParseCityName('%v'): FAILED, expected an error but got the value '%v'", item.input, result)
			} else {
				t.Logf("ParseCityName('%v'): PASSED, expected an error and got an error '%v'", item.input, err)
			}
		} else {
			// expected a value
			if result != item.result {
				t.Errorf("ParseCityName('%v'): FAILED, expected '%v' but got value '%v'", item.input, item.result, result)
			} else {
				t.Logf("ParseCityName('%v'): PASSED", item.input)
			}
		}
	}
}
