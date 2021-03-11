package main

import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()

	Apis(r)
	RunServer(":8080", r)
}
