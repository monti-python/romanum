package main

import (
	"fmt"
	"os"
	"strconv"
	"romanum/pkg/converter"
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

	roman, err := converter.ToRoman(number)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The Roman numeral for %d is %s\n", number, roman)
	// fmt.Println("Halted execution")
}
