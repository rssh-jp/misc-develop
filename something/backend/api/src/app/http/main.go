package main

import (
	"fmt"
    "log"
	"net/http"
	"os"
)

func main() {
    log.Println("START")
    defer log.Println("END")

    err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
        log.Printf("Request: %+v", r)
        fmt.Println("Request Method:", r.Method)
        fmt.Println("Request URL:", r.URL)
        fmt.Println("Request Headers:")
        for key, values := range r.Header {
            fmt.Printf("  %s: %s\n", key, values)
        }
	}))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

