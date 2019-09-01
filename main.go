package main

import (
	"flag"
	"fmt"
	"github.com/climber73/tendermint-challenge/worldx"
	"os"
)

func main() {
	n := flag.Int("n", 2, "number of aliens")
	path := flag.String("path", "", "path to map file")
	flag.Parse()

	if len(*path) == 0 {
		panic("empty path")
	}

	file, err := os.Open(*path)
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
