package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduardor2m/db-metrics/src/internal/models"
	"github.com/eduardor2m/db-metrics/src/internal/repositories/postgres"
	"github.com/google/uuid"
)

type Response struct {
	Message string `json:"message"`
}

type UserDTO struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Preferences []string `json:"preferences"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	userPostgresRepository := postgres.UserPostgresRepository{}

	id := uuid.New()
	newUserModel := models.NewUser(id, user.Name, user.Email, user.Preferences)
	fmt.Println(*newUserModel)
	response := userPostgresRepository.Create(*newUserModel)

	// response := Response{Message: "User created"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "User retrieved"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
