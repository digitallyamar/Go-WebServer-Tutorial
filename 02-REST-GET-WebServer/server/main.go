package main

import (
	"GO-WebServer-Tutorial/02-REST-GET-WebServer/countryCapitals"
	"fmt"
	"net/http"
	"strings"
)

const port = ":9000"

func home(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")

	if urlPathElements[1] == "capital" {
		capital := countryCapitals.Capitals[urlPathElements[2]]

		//If not match found, empty string is returned. Use it for validation!
		if capital != "" {
			fmt.Fprintf(w, capital)
		} else {
			fmt.Fprintf(w, "Sorry! No capital found for this country.")
		}
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)
}
