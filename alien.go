package main

import (
	"fmt"
)

type Alien struct {
	id       int
	cityName string
	steps    int
}

func (a Alien) String() string {
	return fmt.Sprintf("alien{%v,%v,%v}", a.id, a.cityName, a.steps)
}
