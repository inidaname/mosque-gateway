package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	cfg "github.com/inidaname/mosque/api_gateway/config"
	customMiddleware "github.com/inidaname/mosque/api_gateway/middleware"
	"github.com/inidaname/mosque/api_gateway/pkg/utils"
)

// SetupRouter configures and returns the API router
func SetupRouter() *chi.Mux {
	// Create a new router
	router := chi.NewRouter()

	// Add middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(customMiddleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:8080")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		utils.SendResponse(w, http.StatusOK, "API Gateway is running", map[string]any{})
	})

	router.Route("/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			utils.SendResponse(w, http.StatusOK, "API Gateway is running", map[string]any{})
		})
		AuthRoutes(r)
		MosqueRoutes(r)
	})

	return router
}
