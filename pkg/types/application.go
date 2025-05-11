package types

import (
	"log/slog"

	// "github.com/inidaname/mosque/api_gateway/pkg/messages/mailer"
	cache "github.com/inidaname/mosque/api_gateway/pkg/store/cache"
)

type Application struct {
	Config        Config
	Logger        *slog.Logger
	Authenticator Authenticator
	// Mailer        mailer.Mailer
	Cache cache.CacheService
}
