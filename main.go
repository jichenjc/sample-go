package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello, xx")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Handling HTTP requests on %s.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
