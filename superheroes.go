package superheroes

import (
	"math/rand"
	"time"
)

// All returns all superheroes from the list.go
func All() []string {
	return superheroes
}

// Random returns a single superhero from the list.go
func Random() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(superheroes) - 1
	return superheroes[rand.Intn(max-min+1)+min]
}
