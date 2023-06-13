// Dup2 prints the count and text of lines that appear more than once
// in the print. It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, true)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, false)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// A map is a reference to the data structure created by make. When a map is passed to a function
// the function receives a copy of the reference, so any changes the called function make to the underlying
// data structure will be visible through the caller's map reference too.

func countLines(f *os.File, counts map[string]int, readStdin bool) {
	if readStdin {
		fmt.Println("Type 'exit' to stop the program")
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if readStdin && text == "exit" {
			break
		}
		counts[text]++
		if !readStdin && counts[text] > 1 {
			fmt.Printf("File's name in which duplicated line occurs: %s, word duplicated: %s\n", f.Name(), text)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
