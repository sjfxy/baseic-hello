package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}
func rollTwo(size1, size2 int) (int, int) {
	return dieRoll(size1), dieRoll(size2)
}
func returnsNamed(input1 string, input2 int) (theResult string, err error) {
	theResult = "modified " + input1 + ", " + strconv.Itoa(input2)
	return theResult, err
}
func main() {
	fmt.Printf("Rolled a die of size %d, result: %d\n", 6, dieRoll(6))
	res1, res2 := rollTwo(6, 10)

	fmt.Printf("Roolled a pair of dice (%d,%d), results: %d, %d\n", 6, 10, res1, res2)

	named, err := returnsNamed("globule", 42)

	fmt.Printf("Name params returned: '%s', %v\n", named, err)

	var p = person{}

	var p2 = person{"bob", 21}

	var p3 = person{name: "bob", age: 22}

	var p4 = &person{}

	var p5 = &person{"sin", 22}

	var p6 = &person{name: "sin", age: 21}

	fmt.Println(p, p2, p3, p4, p5, p6)

}
