package main

import "fmt"

// package level
const boilingF = 212.0

func main() {
	// var f and c are local to function main
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
	// Output:
	// boiling point = 212F or 100C
}
