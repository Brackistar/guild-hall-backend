package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Brackistar/guild-hall-backend/constants"
	"github.com/Brackistar/guild-hall-backend/controllers"
	"github.com/Brackistar/guild-hall-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Inicio del servidor
func main() {
	// Iniciar libreria de configuración Viper
	configViper()

	serverConfig := new(models.ServerConfig)

	setConfiguration(serverConfig)

	router := gin.Default()

	configureEndpoints(router)

	url := fmt.Sprintf("%s:%d", serverConfig.HostName, serverConfig.Port)
	router.Run(url)
}

func configureEndpoints(router *gin.Engine) {
	router.GET(constants.StatusEndpoint, controllers.TestConnection)
}

// Configura lectura de archivo de configuración con libreria Viper
func configViper() {
	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(constants.Json)
	viper.AddConfigPath(constants.ConfigFilePath)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func setConfiguration(config *models.ServerConfig) {
	config.HostName = viper.GetString(constants.ServerHostConstant)
	config.Port = viper.GetInt(constants.ServerPortConstant)
	config.HasSSL = viper.GetBool(constants.ServerHasSslConstant)
}
