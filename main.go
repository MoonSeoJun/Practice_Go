package main

import (
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/controllers"
	"github.com/MoonSeoJun/FreeWebMaking/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	models.ConnectDB()
	Apis(r)
	RunServer(":8080", r)
}

// RunServer - using https
func RunServer(port string, r *mux.Router) {
	http.ListenAndServe(port, r)
}

// Apis - collected apis
func Apis(r *mux.Router) {

	r.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/getUser/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/getAllUsers", controllers.GetAllUser).Methods("GET")
	r.HandleFunc("/deleteUser/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/updateUser/{id}", controllers.UpdateUser).Methods("PATCH")
}
