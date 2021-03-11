package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MoonSeoJun/FreeWebMaking/models"
)

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
