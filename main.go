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

func determineAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {

	_, err := determineAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handlerFunc)

	//http.ListenAndServe(":3000", nil)//local machine
	//http.ListenAndServe(port, nil)
}
