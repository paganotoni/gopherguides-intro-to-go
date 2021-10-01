package main

import "fmt"

func main() {
	// First time using multiple assignment expression
	// used short names given the short scope of the variables
	s, i, b := "Go", 42, true

	// Discovered the index notation! Love it.
	fmt.Printf("Printing %[1]T(%[1]s)!\n", s)
	fmt.Printf("Printing %[1]T(%[1]d)!\n", i)
	fmt.Printf("Printing %[1]T(%[1]t)!\n", b)
}
