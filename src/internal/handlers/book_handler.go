package handlers

import (
	"encoding/json"
	"net/http"
)

type BookDTO struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Genre    string `json:"genre"`
	Synopsis string `json:"synopsis"`
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book BookDTO
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{Message: "Book created"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Book retrieved"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
