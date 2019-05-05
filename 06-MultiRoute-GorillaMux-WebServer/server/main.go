package main

import (
	"GO-WebServer-Tutorial/02-REST-GET-WebServer/countryCapitals"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
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
			//Returns 404, Not found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Sorry, Resource Not Found!"))
		}
	} else {
		//Returns 400, Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Sorry, Bad API Request!"))
	}
}

func getServerTime(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, time.Now().String())
}

func getRandomNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Int())
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User belongs to category: %v\n", vars["category"])
	fmt.Fprintf(w, "User ID is: %v\n", vars["id"])
}

func main() {
	//Introducing NewServeMux for multi-routing
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", home)
	rtr.HandleFunc("/servertime", getServerTime)
	rtr.HandleFunc("/random", getRandomNumber)
	rtr.HandleFunc("/users/{category}/{id:[0-9]+}", handleUsers)
	log.Fatal(http.ListenAndServe(port, rtr))
}
