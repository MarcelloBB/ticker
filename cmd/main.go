package main

import (
	"github.com/MarcelloBB/ticker/internal/config"
	"github.com/MarcelloBB/ticker/internal/router"
)

// @title           Ticker Uptime API
// @version         1.0
// @description     Uptime monitoring service
// @host            localhost:8080
// @BasePath        /
func main() {
	config.InitializeConfig()
	router.InitRouter()
}
