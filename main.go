package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Brackistar/guild-hall-backend/models"
	"github.com/gin-gonic/gin"
)

// Inicio del servidor
func main() {
	var serverConfig models.ServerConfig
	getConfig(&serverConfig)

	router := gin.Default()

	url := fmt.Sprintf("%s:%d", serverConfig.HostName, serverConfig.Port)
	router.Run(url)
}

// Leer configuración desde archivo de config
func getConfig(config *models.ServerConfig) {
	file, err := os.Open("./Config/serverConfig.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		panic(err)
	}
}
