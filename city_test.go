package main

import (
	"testing"
)

type TestDataItem struct {
	receiver City
	result   []Direction
}

func TestAvailableDirs(t *testing.T) {
	some := &City{}
	testDataItems := []TestDataItem{
		{City{name: "Foo"}, []Direction{}},
		{City{name: "Foo", north: some}, []Direction{North}},
		{City{name: "Foo", south: some}, []Direction{South}},
		{City{name: "Foo", west: some}, []Direction{West}},
		{City{name: "Foo", east: some}, []Direction{East}},
		{City{name: "Foo", north: some, south: some, west: some, east: some}, []Direction{North, South, West, East}},
	}
	for _, item := range testDataItems {
		result := item.receiver.AvailableDirs()
		if testEq(result, item.result) {
			t.Logf("%v.AvailableDirs(): PASSED", item.receiver)
		} else {
			t.Errorf("%v.AvailableDirs(): FAILED, expected %v but got %v", item.receiver, item.result, result)
		}
	}
}

func testEq(a, b []Direction) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
