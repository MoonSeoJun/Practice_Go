package apis

import (
	"github.com/MoonSeoJun/FreeWebMaking/controllers"
	"github.com/gorilla/mux"
)

// Apis - collected apis
func Apis(r *mux.Router) {

	r.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/getUser/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/getAllUsers", controllers.GetAllUser).Methods("GET")
	r.HandleFunc("/deleteUser/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/updateUser/{id}", controllers.UpdateUser).Methods("PATCH")
}
