package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	_ "github.com/souravagrawal29/go-playground/bookstore/pkg/config"
	"github.com/souravagrawal29/go-playground/bookstore/pkg/routes"
)

func main() {
	
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}