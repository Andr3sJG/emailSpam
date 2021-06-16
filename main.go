package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello there <h1>")
}

func main() {
	fmt.Println("hi onichan")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
