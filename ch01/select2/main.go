// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
// ref: https://dev.to/chseki/the-select-keyword-in-golang-1fj4
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// $ time go run main.go

var tableNumbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

const (
	_exit = "exit"
)

type message struct {
	text    string
	number  int
	isValid bool
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	msg := make(chan message)
	exit := make(chan bool)

	go readInput(msg, exit)

	for {
		// We’ll use select to await both of these values simultaneously
		select {
		case m := <-msg:
			validateInputAndGetNumber(&m)
			log.Printf("%#v\n", m)
			printMultiplicationTable(m)
			if m.isValid {
				fmt.Print("Give me a number or exit: ")
			} else {
				fmt.Print("Give me a valid number or exit: ")
			}
		case <-exit:
			log.Println("Exiting program...")
			os.Exit(1)
		}
	}

}

func readInput(msg chan<- message, exit chan<- bool) {
	fmt.Print("Give me a number or exit: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if _exit == text {
			exit <- true
		} else {
			msg <- message{
				text: text,
			}
		}
	}
}

func validateInputAndGetNumber(m *message) {
	n, err := strconv.Atoi(m.text)
	if err != nil {
		m.isValid = false
	} else {
		m.number = n
		m.isValid = true
	}
}

func printMultiplicationTable(m message) {
	if m.isValid {
		for _, num := range tableNumbers {
			fmt.Printf("%3d   x %3d   = %3d\n", m.number, num, m.number*num)
		}
	}
}
