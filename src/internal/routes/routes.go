package routes

import (
	"net/http"

	"github.com/eduardor2m/db-metrics/src/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/user", handlers.CreateUserHandler)
	r.Get("/user", handlers.GetUserHandler)

	r.Post("/book", handlers.CreateBookHandler)
	r.Get("/book", handlers.GetBookHandler)

	r.Post("/review", handlers.CreateReviewHandler)
	r.Get("/review", handlers.GetReviewHandler)

	return r
}
