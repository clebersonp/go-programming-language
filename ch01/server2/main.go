// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	// the program would have a serious bug called a 'race condition'
	// when two or more goroutine changes the count variable at the same time.
	// To avoid this problem, we must ensure that at most one goroutine
	// accesses the variable at a time, which is the purpose of the mu.Lock() and mu.Unlock() calls
	// that bracket each access of count
	mu.Lock()
	count++
	mu.Unlock()
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
}

// counter echoes
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
	if err != nil {
		log.Fatal(err)
	}
}
