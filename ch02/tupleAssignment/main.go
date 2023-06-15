package main

import "fmt"

func main() {
	// tuple assignment allows several variables to be assigned at once
	var x, y int // declare all at once to the same int type
	x, y = 15, 5 // assignment all at once
	fmt.Printf("x = %v, y = %v\n", x, y)

	swap(x, y)
	fmt.Printf("GCD(%v, %v) = %v\n", x, y, gcd(x, y))
	fmt.Printf("Fib(10) = %v\n", fib(10))

	// Map lookup, type assertion, and channel receive can return bool value to evaluate it
	// v, ok = m[key]	// map lookup
	// v, ok = x.(T)	// type assertion
	// v, ok = <-ch		// channel receive

}

// gcd Greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		// All of right-hand side expressions are evaluated before any of the variables are updated
		x, y = y, x%y
	}
	return x
}

func swap(x, y int) {
	x, y = y, x // changes only local, not affect outside
	fmt.Printf("Swap: x = %v, y = %v\n", x, y)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
