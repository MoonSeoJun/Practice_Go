package main

import (
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/controllers"
	"github.com/gorilla/mux"
)

// Apis sort
func Apis(r *mux.Router) {

	r.HandleFunc("/method", controllers.JSONMarshal).Methods("POST")
}

// RunServer - using https
func RunServer(port string, r *mux.Router) {
	http.ListenAndServe(port, r)
}
