package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/models"
	"github.com/gorilla/mux"
)

// CreateUser - save user info
func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		panic(err)
	}

	createUser := models.User{Name: u.Name, Age: u.Age, Gender: u.Gender}

	models.DB.Create(&createUser)

	userJSON, err := json.Marshal(u)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userJSON)
}

// GetUser - get user info
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)
	id := params["id"]

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		panic(err)
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

// GetAllUser - get all users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	models.DB.Find(&users)

	userJSON, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
