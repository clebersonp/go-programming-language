package main

import "fmt"

func main() {

	// nil is a predeclared identifier representing the zero value for a
	// pointer, channel, func, interface, map, or slice type.
	// Type must be a pointer, channel, func, interface, map, or slice type
	// Zero value for numbers is Zero(0), strings is "", and boolean is false

	var x, y int
	fmt.Printf("&x == &x? %t\n&x == &y? %t\n&x == nil? %t\n&y == nil? %t\nx == 0? %t\ny == 0? %t\n",
		&x == &x, &x == &y, &x == nil, &y == nil, x == 0, y == 0)
}
