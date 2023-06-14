// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
// ref: https://gobyexample.com/select
package main

import (
	"log"
	"time"
)

// $ time go run main.go
func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	c1 := make(chan string)
	c2 := make(chan string)

	log.Println("Start...")

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// We’ll use select to await both of these values simultaneously, printing each one as it arrives
		select {
		case msg1 := <-c1:
			log.Printf("received %s from iteration %d", msg1, i)
		case msg2 := <-c2:
			log.Printf("received %s from iteration %d", msg2, i)
		}
	}

	log.Println("Ended")
}
