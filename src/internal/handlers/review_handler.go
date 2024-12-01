package handlers

import (
	"encoding/json"
	"net/http"
)

type ReviewDTO struct {
	UserID  string `json:"user_id"`
	BookID  string `json:"book_id"`
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
}

func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review ReviewDTO
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{Message: "Review created"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Review retrieved"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
