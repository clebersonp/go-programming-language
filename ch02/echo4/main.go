package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// examples to run:
// $ go run main.go -n=false -s - this is just a test
// $ go run main.go -n=true -s - this is just a test
// $ go run main.go -n -s / this is just a test
// $ go run main.go -s , this is just a test
// $ go run main.go this is just a test
// for help of usage:
// $ go run main.go -h
// $ go run main.go --h
// $ go run main.go -help
// $ go run main.go --help

// flag doc: $ go doc flag

func main() {
	// Parse parses the command-line flags from os.Args[1:].
	// Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	fmt.Printf("Value of 'n' flag: %v\n", *n)
	fmt.Printf("Value of 'sep' flag: %v\n", *sep)
	fmt.Printf("Flag args values: %v\n\n", flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))

	// newline if 'n' flag is false
	if !*n {
		fmt.Println()
	}
}
