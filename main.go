package main

import (
	api "github.com/inidaname/mosque/api_gateway/cmd/server"
	"github.com/inidaname/mosque/api_gateway/config"
)

func main() {

	// // Load configuration
	// cfg := config.LoadConfig()
	app := config.CreateApplication()
	app.Logger.Error("server error", "result", api.Run())
}
