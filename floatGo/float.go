package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&input)
	truncated := int(math.Trunc(input))
	fmt.Printf("Truncated integer of %v is %v\n", input, truncated)
}
