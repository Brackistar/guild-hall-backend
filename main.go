package main

import (
	"fmt"

	"os"

	"github.com/Brackistar/guild-hall-backend/constants"
	"github.com/Brackistar/guild-hall-backend/controllers"
	"github.com/Brackistar/guild-hall-backend/helpers"
	"github.com/Brackistar/guild-hall-backend/interfaces"
	"github.com/Brackistar/guild-hall-backend/models"
	"github.com/Brackistar/guild-hall-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var loggerHelper interfaces.ILogHelper
var logger *logrus.Logger

// Inicio del servidor
func main() {
	loggerHelper = helpers.NewLogrusHelper()
	logger = loggerHelper.ConfigureLogger()

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
	configureHomeController(router)
	configureAdventurerController(router)
}

func configureHomeController(router *gin.Engine) {
	homeController := controllers.NewHomeController(logger)
	router.GET(constants.StatusEndpoint, homeController.TestConnection)
}

func configureAdventurerController(router *gin.Engine) {
	service := services.NewMockAdventurerService()
	adventurerController := controllers.NewAdventurerController(*service, logger)

	router.GET(constants.GetAdventurerEndpoint, adventurerController.GetAdventurer)
}

// Configura lectura de archivo de configuración con libreria Viper
func configViper() {
	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(constants.Json)
	viper.AddConfigPath(constants.ConfigFilePath)

	err := viper.ReadInConfig()

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func setConfiguration(config *models.ServerConfig) {
	config.HostName = viper.GetString(constants.ServerHostConstant)
	config.Port = viper.GetInt(constants.ServerPortConstant)
	config.HasSSL = viper.GetBool(constants.ServerHasSslConstant)
}
