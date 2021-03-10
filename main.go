package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sayHello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	}).Methods("GET")
	r.HandleFunc("/method", JSONMarshal).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// JSONMarshal is marshaling struct
func JSONMarshal(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		panic(err)
	}

	userJSON, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
