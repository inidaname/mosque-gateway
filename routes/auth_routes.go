package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inidaname/mosque/api_gateway/client"
	"github.com/inidaname/mosque/api_gateway/handlers"
)

func AuthRoutes(r chi.Router) http.Handler {
	authClient := client.NewAuthClient("localhost:50051")
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handlers.RegisterUser(authClient))
		r.Post("/login", handlers.LoginUser(authClient))
		r.Route("/forgot-password", func(r chi.Router) {
			r.Post("/", handlers.ForgotPassword(authClient))
			r.Get("/validate-token/{token}", handlers.ValidatePasswordToken(authClient))
		})
	})

	return r
}
