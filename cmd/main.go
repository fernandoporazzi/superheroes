package main

import (
	"fmt"

	"github.com/fernandoporazzi/superheroes"
)

func main() {
	fmt.Println(superheroes.Random())
	fmt.Println(superheroes.All())
}
