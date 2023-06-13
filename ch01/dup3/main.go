// Reads all content of given files as argument and count the lines 'words'
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("dup3: %v\n", err)
			continue // goes to the next iteration
		}
		for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
