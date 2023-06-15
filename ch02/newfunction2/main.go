package main

import "fmt"

func main() {

	// both creation of pointer to *int are equivalents
	p := newIntUnnamedVariable()
	q := newIntDummyLocalVariable()
	fmt.Printf("'p' pointer address: %v\n'q' pointer address: %v\n", p, q)
	fmt.Printf("'p' pointer address == 'q' pointer address? %t\n", p == q)
}

func newIntUnnamedVariable() *int {
	return new(int)
}

func newIntDummyLocalVariable() *int {
	var dummy int
	return &dummy
}
