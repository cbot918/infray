package main

import (
	"fmt"
	"net/http"
)

const port = ":8081"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "auth service /")
	})
	fmt.Printf("auth service listening: %s", port)
	http.ListenAndServe(port, nil)
}
