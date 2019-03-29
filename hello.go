package main

import (
	"fmt"
	"math"
	"math/rand"
)

var i, j int = 9, 10

func add(x, y int) int {
	return x + y
}

func main() {
	// Able to name variables without the var word. Just use :=
	n := 18
	str := "Ned"
	sum := 0
	// Must import what package you need to use after the package main
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Println("The square root of 7 is", math.Sqrt(7))
	fmt.Printf("Now you have %g problems./n", math.Sqrt(8))
	fmt.Println("My name is ", str)

	// Functions can be declared outside of the func main and then used inside. Only code inside the func main will be executed
	fmt.Println(add(i, n))
	fmt.Println(add(4, 7))
	fmt.Println(add(i, j))

	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
