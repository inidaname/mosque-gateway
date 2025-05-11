package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/inidaname/mosque/api_gateway/client"
	"github.com/inidaname/mosque/api_gateway/handlers"
)

func MosqueRoutes(r chi.Router) http.Handler {
	mosqueClient := client.NewMosqueClient("localhost:50052")
	r.Route("/mosque", func(r chi.Router) {
		r.Post("/", handlers.CreateMosque(mosqueClient))
		r.Get("/", handlers.ListMosque(mosqueClient))
		r.Put("/", handlers.UpdateMosque(mosqueClient))
	})

	return r

}
