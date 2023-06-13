package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a new goroutine
	}

	// this runs in the main goroutine and receive values from others goroutines by channel
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%3dms elapsed\n", time.Since(start).Milliseconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch any error
		return                // return if there's any error
	}

	nBytes, err := io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	millis := time.Since(start).Milliseconds()
	ch <- fmt.Sprintf("%3dms\t%7d\t%s", millis, nBytes, url)
}
