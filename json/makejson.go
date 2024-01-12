package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter address: ")
	address, _ := reader.ReadString('\n')

	// Trim the newline character from the input
	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
