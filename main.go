package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")

	// Scan the input
	scanner.Scan()
	FirstName := scanner.Text()
	fmt.Print("Enter your name: ")

	// Scan the input
	scanner.Scan()
	LastName := scanner.Text()

	fmt.Printf("Welcome to Go, %v %v\n", FirstName, LastName)
}
