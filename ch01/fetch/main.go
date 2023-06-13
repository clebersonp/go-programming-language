// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	prefix      = "https://"
	json        = "application/json"
	html        = "text/html"
	contentType = "Content-Type"
)

func main() {
	for _, url := range os.Args[1:] {
		if hasPrefix := strings.HasPrefix(url, prefix); !hasPrefix {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		//b, err := io.ReadAll(resp.Body)
		//fmt.Println(resp.Header[contentType])
		extension := ".txt"
		contents, ok := resp.Header[contentType]
		if ok {
			for _, value := range contents {
				if strings.Contains(value, json) {
					extension = ".json"
					break
				} else if strings.Contains(value, html) {
					extension = ".html"
					break
				}
			}
		}
		file, err := os.Create(strconv.FormatInt(time.Now().UnixMilli(), 10) + extension)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(file, resp.Body)
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Response status code: %s\n", resp.Status)
	}
}
