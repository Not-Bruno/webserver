package main

import (
	"fmt"
	"log"
	"net/http"
)

var gPort string = "8080"

func main() {
	// run index.html on http://localhost:8080
	fileServer := http.FileServer(http.Dir("./app"))
	http.Handle("/", fileServer)

	// Call homeHandler Function in HandleFunc
	http.HandleFunc("/state", stateHandler)

	fmt.Println("Starting Webserver on Port: " + gPort)
	// Starting Webserver and Check for startingerror
	if err := http.ListenAndServe(":"+gPort, nil); err != nil {
		log.Fatal(err)
	}
}

// Handle request on http://localhost:8008/home
func stateHandler(w http.ResponseWriter, r *http.Request) {
	// if sub site is not localhost/home
	if r.URL.Path != "/state" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method "+r.Method+" not Supportet", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Webserver started!\nListening on Port: "+gPort)
}
