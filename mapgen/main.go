package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/climber73/tendermint-challenge/worldx"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	n := flag.Int("n", 3, "number of rows")
	m := flag.Int("m", 3, "number of cols")
	path := flag.String("path", "", "path to map file")
	flag.Parse()

	if len(*path) == 0 {
		exit(fmt.Errorf("empty path"))
	}

	file, err := createFile(*path)
	if err != nil {
		exit(err)
	}
	if file == nil {
		os.Exit(0)
	}
	defer file.Close()

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

	for _, city := range cities {
		s := fmt.Sprintf("%v\n", *city)
		if _, err := file.WriteString(s); err != nil {
			exit(err)
		}
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

func createFile(path string) (*os.File, error) {
	info, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			return nil, fmt.Errorf("'%s' is a directory.", path)
		}
		question := fmt.Sprintf("File '%s' already exists. Do you want to overwrite it? (yes/no)", path)
		if ask(question) {
			return os.Create(path)
		} else {
			return nil, nil
		}
	}
	if err := ensureDir(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}
	return os.Create(path)
}

func ask(msg string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(msg)
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)
		switch strings.ToLower(answer) {
		case "yes":
			return true
		case "no":
			return false
		}
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(-1)
}

func ensureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}
