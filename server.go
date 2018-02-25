package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world \n")
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
