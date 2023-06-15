package main

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, point to x variable
	fmt.Println(&x) // address of x variable
	fmt.Println(p)  // p points to address of x variable
	fmt.Println(x)  // x variable value
	fmt.Println(*p) // p points to x variable value
	fmt.Printf("Address of 'x' == 'p' points to int? %t\n", &x == p)
	*p = 2 // change the value of x variable that p points to
	fmt.Println(x)
	fmt.Println(*p)

}
