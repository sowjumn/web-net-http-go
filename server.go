package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "home.html")
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
