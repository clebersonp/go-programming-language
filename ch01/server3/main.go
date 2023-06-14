package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "%s %s %s %s\n", r.Method, r.URL, r.URL.Path, r.Proto)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range r.Header {
		_, err = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = fmt.Fprintf(w, "Host = %q\n", r.Host)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	if err = r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, err = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		if err != nil {
			log.Fatal(err)
		}
	}
}
