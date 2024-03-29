package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanln(&a)

	fmt.Print("Enter initial velocity: ")
	fmt.Scanln(&vo)

	fmt.Print("Enter initial displacement: ")
	fmt.Scanln(&so)

	fmt.Print("Enter time: ")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, vo, so)
	displacement := fn(t)

	fmt.Println("Displacement:", displacement)
}
