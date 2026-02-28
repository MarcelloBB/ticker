package main

import (
	"github.com/MarcelloBB/ticker/internal/config"
	"github.com/MarcelloBB/ticker/internal/router"
)

func main() {
	config.InitializeConfig()
	router.InitRouter()
}
