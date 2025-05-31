package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from [user]!")
	})

	port := "8080"
	fmt.Printf("Starting [service user] on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
