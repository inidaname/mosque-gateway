package config

import (
	"log"
	"time"

	"github.com/inidaname/mosque/api_gateway/pkg/types"
	"github.com/joho/godotenv"
)

// Global Config

// LoadConfig loads environment variables and sets the configuration
func LoadConfig() types.Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load .env file")
	}

	appConfig := types.Config{
		Addr:   GetString("ADDR", ":8080"),
		Env:    GetString("ENV", "development"),
		ApiURL: GetString("EXTERNAL_URL", "localhost:8080"),
		Auth: types.AuthConfig{
			Basic: types.BasicConfig{
				User: GetString("AUTH_BASIC_USER", "admin"),
				Pass: GetString("AUTH_BASIC_PASS", "admin"),
			},
			Token: types.TokenConfig{
				Secret: GetString("AUTH_TOKEN_SECRET", "supersecret"),
				Exp:    time.Hour * 24,
				Iss:    GetString("AUTH_TOKEN_NAME", "mosques"),
			},
		},
		Mailer: types.MailConfig{
			FromEmail: GetString("FROM_EMAIL", "support@mainheart.co"),
			Resend: types.ResendConfig{
				ApiKey: GetString("RESEND_API_KEY", ""),
			},
		},
		Dospace: types.DigitalOceanSpace{
			AccessKey: GetString("DO_SPACE_ACCESS_KEY", ""),
			SecretKey: GetString("DO_SPACE_SECRET_KEY", ""),
			SpaceName: GetString("DO_SPACE_BUCKET_NAME", "finecore"),
			Endpoint:  GetString("DO_SPACE_ENDPOINT", "https://us-east-1.digitaloceanspaces.com"),
		},
	}
	return appConfig
}
