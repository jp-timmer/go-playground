package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers []int
	numbers = make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter an integer (or type 'X' to exit): ")
		fmt.Scan(&input)
		num, err := strconv.Atoi(input)
		if input == "X" {
			break
		} else if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		numbers = append(numbers, num)
		sort.Ints(numbers)

		fmt.Println("Sorted slice:", numbers)
	}
}
