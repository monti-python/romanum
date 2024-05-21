package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("Starting up romanum server...")

	if len(os.Args) < 2 {
		fmt.Println("Usage: romanum <number>")
		os.Exit(1)
	}

	input := os.Args[1]
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Error: Invalid number '%s'\n", input)
		os.Exit(1)
	}

	fmt.Printf("Input number is %d\n", number)
	// fmt.Println("Halted execution")
}
