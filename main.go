package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/climber73/tendermint-challenge/worldx"
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

	m, err := worldx.NewMap(file)
	if err != nil {
		panic(err)
	}

	r := worldx.Randomizer{}
	world, err := worldx.NewWorld(m, *n, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Initial world:\n%v\n", world)

	if err := world.Run(); err != nil {
		panic(err)
	}

	fmt.Printf("The rest of the world:\n%v\n", world)

}
