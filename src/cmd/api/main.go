package main

import (
	"net/http"

	"github.com/eduardor2m/db-metrics/src/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", routes.RegisterRoutes())
}
