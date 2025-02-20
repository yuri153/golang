package controllers

import (
	"encoding/json"
	"fmt"
	"golang/api-go-rest/database"
	"golang/api-go-rest/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality

	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personalityId := vars["id"]

	var personality models.Personality

	database.DB.First(&personality, personalityId)

	json.NewEncoder(w).Encode(personality)
}
