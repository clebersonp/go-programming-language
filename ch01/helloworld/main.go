//Prints the hello world
package main

import "fmt"

func main() {
	// to format and update the go file by standard go tools: $ gofmt -w filename
	// or we can substitute with the standard go tools: $ goimports -w filename
	// go imports format the code and auto import packages
	// by default, Golang use UTF-8 encoding for string
	fmt.Println("Hello, 世界")
}
