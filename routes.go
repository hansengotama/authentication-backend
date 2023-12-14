package main

import (
	"github.com/hansengotama/authentication-backend/internal/handler/requestotphandler"
	"github.com/hansengotama/authentication-backend/internal/handler/validateotphandler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth-otp", func(r chi.Router) {
		r.Post("/request", requestotphandler.HandleRequestOTPAuth)
		r.Post("/validate", validateotphandler.HandleValidateOTPAuth)
	})

	return r
}
