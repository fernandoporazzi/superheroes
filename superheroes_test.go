package superheroes

import "testing"

func TestAll(t *testing.T) {
	superheroes := All()

	if len(superheroes) == 0 {
		t.Errorf("Expect list to not be empty")
	}
}

func TestRandom(t *testing.T) {
	randomSuperhero := Random()

	if randomSuperhero == "" {
		t.Errorf("Expect random superhero to exist")
	}
}
