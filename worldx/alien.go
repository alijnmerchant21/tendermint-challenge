package worldx

import (
	"fmt"
)

type Alien struct {
	id       int
	cityName string
	steps    int
}

func (a Alien) String() string {
	return fmt.Sprintf("alien{#%v, city=%v, steps=%v}", a.id, a.cityName, a.steps)
}
