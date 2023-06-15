package main

import "fmt"

func main() {
	// p, of type *int, points to an unnamed int variable with Zero value(0)
	p := new(int) // new function create a pointer to type of T with zero value
	fmt.Printf("Address and value that 'p' unnamed variable of type int points to: %v, %v\n", p, *p)
	*p = 2
	fmt.Printf("Value of 'p' after changes: %v\n", *p)
}
