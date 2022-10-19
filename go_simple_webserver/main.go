package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello_handler(response http.ResponseWriter, request *http.Request) {
	// If the user made a request to a path that doesn't exist
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}

	// If the method wasn't GET
	if request.Method != "GET" {
		http.Error(response, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(response, "Hello")
}

func form_handler(response http.ResponseWriter, request *http.Request) {
	// In case something goes wrong while parsing the form
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(response, "POST request successful\n")

	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(response, "Name: %s\nAddress: %s", name, address)
}

func main() {
	file_server := http.FileServer(http.Dir("./static"))

	http.Handle("/", file_server)
	http.HandleFunc("/form", form_handler)
	http.HandleFunc("/hello", hello_handler)

	fmt.Println("Starting server at port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
