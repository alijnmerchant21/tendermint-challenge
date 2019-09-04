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
		exit(fmt.Errorf("empty path"))
	}

	file, err := os.Open(*path)
	if err != nil {
		exit(err)
	}
	defer file.Close()

	m, err := worldx.NewMap(file)
	if err != nil {
		exit(err)
	}

	r := worldx.Randomizer{}
	world, err := worldx.NewWorld(m, *n, r)
	if err != nil {
		exit(err)
	}

	if err := world.Run(); err != nil {
		exit(err)
	}

	fmt.Printf("The rest of the world:\n%v\n", world)

}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(-1)
}
