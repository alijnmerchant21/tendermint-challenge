package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestNewWorld(t *testing.T) {
	var w *World
	m := createMap("Foo east=Bar\nBar west=Foo")
	s := rand.NewSource(0)
	r := rand.New(s)

	w, _ = NewWorld(m, 3, r)
	if len(w.aliens) != 2 {
		t.Errorf("NewWorld: FAILED, expected aliens number is 2, but got '%v'\n", len(w.aliens))
	}

	w, _ = NewWorld(m, 1, r)
	if len(w.aliens) != 1 {
		t.Errorf("NewWorld: FAILED, expected aliens number is 1, but got '%v'\n", len(w.aliens))
	}

	t.Logf("%v\n", w.Map.String())
}

func createMap(s string) *Map {
	b := []byte(s)
	reader := bytes.NewReader(b)
	m, _ := NewMap(reader)
	return m
}
