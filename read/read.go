package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	// Prompt the user for the name of the text file
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to store the structs
	var names []name

	// Read each line of the file and create a struct for each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := splitFields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		names = append(names, name{fname: fields[0], lname: fields[1]})
	}

	// Print the first and last names found in each struct
	for _, n := range names {
		fmt.Println("First Name:", n.fname)
		fmt.Println("Last Name:", n.lname)
	}
}

// splitFields splits a line into fields based on a single space separator
func splitFields(line string) []string {
	return strings.Split(line, " ")
}
