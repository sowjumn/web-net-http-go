package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world \n")
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome Home!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
