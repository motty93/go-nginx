package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")

		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":8080", nil)
}
