package config

import (
	"log/slog"
	"os"
	"sync"
	"time"

	cache "github.com/inidaname/mosque/api_gateway/pkg/store/cache"
	// database "github.com/inidaname/mosque/api_gateway/store/database"
	"github.com/inidaname/mosque/api_gateway/pkg/utils"

	"github.com/inidaname/mosque/api_gateway/pkg/types"
)

var (
	instance *types.Application
	once     sync.Once
)

// CreateApplication initializes your application dependencies and ensures only one instance is created
func CreateApplication() *types.Application {
	once.Do(func() {
		cfg := LoadConfig()

		// Initialize JWT authenticator
		jwtAuthenticator := utils.NewJWTAuthenticator(
			cfg.Auth.Token.Secret,
			cfg.Auth.Token.Iss,
			cfg.Auth.Token.Iss,
		)

		// Initialize logger
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		environment := os.Getenv("ENV")
		if environment == "" {
			environment = "development" // default to development if not set
		}

		// var store *db.Queries
		// var dbPool *pgxpool.Pool
		// var err error

		// Initialize thread-safe cache
		cacheService := cache.NewCacheService(5*time.Minute, 10*time.Minute)

		instance = &types.Application{
			Config:        cfg,
			Logger:        logger,
			Authenticator: jwtAuthenticator,
			Cache:         *cacheService,
		}
	})

	return instance
}
