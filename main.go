package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	n := flag.Int("n", 2, "number of aliens")
	flag.Parse()

	// todo: get file path from args

	file, err := os.Open("testdata/map.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m, err := NewMap(file)
	if err != nil {
		panic(err)
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	world, err := NewWorld(m, *n, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Initial world:\n%v\n", world)

	for !world.StopCondition() {
		for id := range world.aliens {
			world.MoveAlien(id)
		}
	}

	fmt.Printf("The rest of the world:\n%v\n", world)

}
