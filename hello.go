package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Println("The square root of 7 is", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.", math.Sqrt(8))
}
