package main

import "fmt"

type power struct {
	attack  int
	defense int
}
type location struct {
	x float32
	y float32
	z float32
}
type noPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

func main() {
	fmt.Println("Structs...")

	demon := noPlayerCharacter{
		name:  "Alires",
		speed: 21,
		hp:    1000,
		power: power{attack: 75, defense: 50},
		loc:   location{x: 1075.123, y: 521.123, z: 211.231},
	}
	fmt.Println(demon)

	anoterDemon := noPlayerCharacter{
		name:  "dsd",
		speed: 0,
		hp:    0,
		power: power{},
		loc:   location{},
	}
	fmt.Println(anoterDemon)
}
