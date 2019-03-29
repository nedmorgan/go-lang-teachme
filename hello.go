package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var i, j int = 9, 10

func add(x, y int) int {
	return x + y
}

func main() {
	// Able to name variables without the var word. Just use :=
	n := 18
	str := "Ned"
	sum := 1
	z := float64(4)
	// Must import what package you need to use after the package main
	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Println("The square root of 7 is", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(8))
	fmt.Println("My name is ", str)
	fmt.Println(z)

	// Functions can be declared outside of the func main and then used inside. Only code inside the func main will be executed
	fmt.Println(add(i, n))
	fmt.Println(add(4, 7))
	fmt.Println(add(i, j))

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Using a switch statement in Go
	fmt.Println("Go is running on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}

	fmt.Println("What time is it now?")
	timeNow := time.Now()
	switch {
	case timeNow.Hour() < 12:
		fmt.Println("Good Morning, the hour is", timeNow.Hour(), "AM")
	case timeNow.Hour() < 17:
		fmt.Println("Good Afternoon, the hour is", (timeNow.Hour() - 12), "PM")
	default:
		fmt.Println("Good night, the time is", (timeNow.Hour() - 12), "PM")
	}
	fmt.Println(timeNow.Minute())
	fmt.Println(timeNow.Hour())
	fmt.Println(timeNow.Month())
}
