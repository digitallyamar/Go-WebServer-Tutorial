package main

import (
	"fmt"
	"net/http"
)

const Port = ":9000"

func Home(w http.ResponseWriter, f *http.Request) {
	response := "Hello, World!"
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(Port, nil)
}