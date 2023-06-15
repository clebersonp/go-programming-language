package main

import "fmt"

func main() {
	var p = f() // variable p is a pointer to int type
	fmt.Printf("p address point to: %v\np value of points to: %v\n", p, *p)

	// each call of f() returns a distinct value:
	firstCall := f()
	secondCall := f()
	fmt.Printf("First call 'f()' address pointer: %v\nSecond call 'f()' address pointer: %v\n",
		firstCall, secondCall)
	fmt.Printf("First call to 'f()' address pointer == second call to 'f()' address pointer? %t\n",
		firstCall == secondCall)
	fmt.Printf("Value of first call pointer: %v\nValue of second call pointer: %v\n",
		*firstCall, *secondCall)
}

// f function return the pointer to int type
func f() *int {
	v := 1
	return &v
}
