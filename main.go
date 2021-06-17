package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello there <h1>")

}

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", handlerFunc)
	log.Print(port)
	http.ListenAndServe(":"+port, nil) //local machine

}
