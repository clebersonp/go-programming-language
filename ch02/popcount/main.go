package main

import "fmt"

// package initialization

// pc[i] is the population count of i.
var pc [256]byte

// init function is called automatically when the program starts, in the order in which they are declared
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	a := uint64(19)
	fmt.Printf("a: %v, popcount: %v\n", a, PopCount(a))
}
