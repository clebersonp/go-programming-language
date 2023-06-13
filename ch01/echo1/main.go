// Echo1 prints its command-line arguments passed by the terminal
package main

import (
	"fmt"
	"os"
)

// e.g: $ go run main.go "first argument" "second argument"

func main() {
	var s, sep string // just declaration, but go will be initialized with Zero-Value. 0 for number and ""(empty) for strings
	// firt index(0) is the program name
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
