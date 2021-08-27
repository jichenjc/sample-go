package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("start main")
	log.Println(runtime.GOARCH)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world! That is arch ")
		fmt.Fprint(w, runtime.GOARCH)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Handling HTTP requests on %s.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
