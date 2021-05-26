package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srinivas/fileserver/controllers"
)

func main() {

	//initialize mux router
	router := mux.NewRouter()

	//handlers for file operations
	router.HandleFunc("/files", controllers.ProcessFile).Methods("POST")
	router.HandleFunc("/statistics", controllers.GetStatistics).Methods("GET")

	//serve app on the specified port
	http.ListenAndServe(":8080", router)

}
