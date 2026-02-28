package router

import (
	"fmt"

	"github.com/MarcelloBB/plata/internal/config"
	"github.com/MarcelloBB/plata/internal/db"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	db.InitRedis()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}

	apiPort := fmt.Sprintf(":%d", config.LoadConfigIni("server", "port", 8080).(int))

	server := gin.Default()
	RegisterRoutes(server, dbConnection)
	server.Run(apiPort)

	return server
}
