package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from [chat]!")
	})

	port := "8080"
	fmt.Printf("Starting [service chat] on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
