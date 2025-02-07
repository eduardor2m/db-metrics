package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eduardor2m/db-metrics/src/internal/models"
	"github.com/eduardor2m/db-metrics/src/internal/repositories/mongodb"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type Response struct {
	Message string `json:"message"`
}

type UserDTO struct {
	Name        string `faker:"username"`
	Email       string `faker:"email"`
	Preferences string `faker:"word"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var usersData [3]UserDTO

	faker.FakeData(&usersData)

	fmt.Println(usersData)

	userPostgresRepository := mongodb.UserMongodbRepository{}

	id := uuid.New()
	newUserModel := models.NewUser(id, usersData[1].Name, usersData[1].Email, usersData[1].Preferences)
	fmt.Println(*newUserModel)
	response := userPostgresRepository.Create(*newUserModel)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// var user UserDTO
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// var usersFakerData [3]UserDTO

	// faker.FakeData(&usersFakerData)

	// fmt.Println(usersFakerData)

	// fmt.Println(user)

	// userPostgresRepository := postgres.UserPostgresRepository{}

	// id := uuid.New()
	// newUserModel := models.NewUser(id, user.Name, user.Email, user.Preferences)
	// fmt.Println(*newUserModel)
	// response := userPostgresRepository.Create(*newUserModel)

	// // response := Response{Message: "User created"}

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "User retrieved"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
