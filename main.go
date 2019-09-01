package main

import (
	"flag"
	"fmt"
	"os"
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

	r := Randomizer{}
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
