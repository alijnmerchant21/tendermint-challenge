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

	fmt.Printf("%v\n", world)

	for !world.StopCondition() {
		fmt.Printf("len of aliens = %v\n", len(world.aliens))
		for id := range world.aliens {
			fmt.Printf("run for %v\n", id)
			world.MoveAlien(id)
		}
		fmt.Printf("%v\n", world)
	}

}
