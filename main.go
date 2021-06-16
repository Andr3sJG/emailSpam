package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello there <h1>")
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println("hi onichan")
	http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe(":3000", nil)//local machine
	http.ListenAndServe(port, nil)
}
