package main

import (
	"fmt"
	"go_book_management_system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("Server running on port 80")
	log.Fatal(http.ListenAndServe(":80", router))
}
