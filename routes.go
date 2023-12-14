package main

import (
	"github.com/hansengotama/authentication-backend/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth-otp", func(r chi.Router) {
		r.Post("/request", handler.RequestOTP)
		r.Post("/validate", handler.ValidateOTP)
	})

	return r
}
