package worldx

import (
	"math/rand"
	"time"
)

type Random interface {
	Perm(int) []int
	Intn(int) int
}

type Randomizer struct {
	seeded bool
}

func (r Randomizer) Intn(n int) int {
	if !r.seeded {
		r.seed()
	}
	num := rand.Intn(n)
	return num
}

func (r Randomizer) Perm(n int) []int {
	if !r.seeded {
		r.seed()
	}
	num := rand.Perm(n)
	return num
}

func (r Randomizer) seed() {
	rand.Seed(time.Now().UnixNano())
	r.seeded = true
}

// for tests:

type FakeRandomizer struct {
	fakeIntn int
}

func (r FakeRandomizer) Intn(n int) int {
	return r.fakeIntn
}

func (r FakeRandomizer) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	return m
}
