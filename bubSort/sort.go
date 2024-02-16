package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	// Prompt the user to enter the array elements
	fmt.Println("Enter the array elements (one element per line, empty line to finish):")

	// Read input lines until an empty line is encountered
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	// Parse the input lines into an integer array
	var arr []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers only.")
			return
		}
		arr = append(arr, num)
	}

	fmt.Println("Before sorting:", arr)
	BubbleSort(arr)
	fmt.Println("After sorting:", arr)
}
