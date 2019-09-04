package main

import (
	"flag"
	"fmt"
	"github.com/climber73/tendermint-challenge/worldx"
	"os"
)

func main() {
	n := flag.Int("n", 3, "number of rows")
	m := flag.Int("m", 3, "number of cols")
	path := flag.String("path", "", "path to map file")
	flag.Parse()

	if len(*path) == 0 {
		panic("empty path")
	}

	cities := make(map[string]*worldx.City, *n**m)
	for i := 0; i < *n; i++ {
		for j := 0; j < *m; j++ {
			name := cityName(i, j)
			cities[name] = worldx.NewCity(
				name,
				northernName(i, j),
				southernName(i, j, *n),
				westernName(i, j),
				easternName(i, j, *m),
			)
		}
	}

	file, err := os.Create(*path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, city := range cities {
		s := fmt.Sprintf("%v\n", *city)
		file.WriteString(s)
	}

	fmt.Printf("map (%vx%v) created\n", *n, *m)
}

func cityName(i, j int) string {
	return fmt.Sprintf("C-%v-%v", i, j)
}

func northernName(i, j int) string {
	if i < 1 {
		return ""
	} else {
		return cityName(i-1, j)
	}
}

func southernName(i, j, n int) string {
	if i >= n-1 {
		return ""
	} else {
		return cityName(i+1, j)
	}
}

func westernName(i, j int) string {
	if j < 1 {
		return ""
	} else {
		return cityName(i, j-1)
	}
}

func easternName(i, j, m int) string {
	if j >= m-1 {
		return ""
	} else {
		return cityName(i, j+1)
	}
}
