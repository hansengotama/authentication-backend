package main

import (
	"github.com/hansengotama/authentication-backend/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth-otp", func(r chi.Router) {
		r.Get("/request", handler.RequestOTP)

		r.Get("/verify", func(writer http.ResponseWriter, request *http.Request) {

		})
	})

	return r
}
