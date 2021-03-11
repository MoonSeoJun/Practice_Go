package main

import (
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/apis"
	"github.com/MoonSeoJun/FreeWebMaking/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	models.ConnectDB()
	apis.Apis(r)

	http.ListenAndServe(":8080", r)
}
