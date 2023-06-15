package main

import "fmt"

func main() {
	v := 1
	fmt.Printf("Address of variable 'v' in main() function: %v\n", &v)
	fmt.Printf("Value of variable 'v' before changes: %v\n", v)
	incr(&v)                                                          // v will be 2
	fmt.Printf("Value of variable 'v' after changes: %v\n", incr(&v)) // v will be 3
}

func incr(p *int) int {
	fmt.Printf("Address of variable 'p' in incr() function that points to an int type: %v\n", p)
	*p++ // increments the value of p points to
	fmt.Printf("Value of variable that 'p' points to after changes: %v\n", *p)
	return *p
}
