package main

import (
	"fmt"

	"os"

	"github.com/Brackistar/guild-hall-backend/constants"
	"github.com/Brackistar/guild-hall-backend/controllers"
	"github.com/Brackistar/guild-hall-backend/enums"
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
var errorHandler interfaces.IErrorHandlerService

// Inicio del servidor
func main() {
	loggerHelper = helpers.NewLogrusHelper()
	logger = loggerHelper.ConfigureLogger()
	errorHandler = services.NewErrorHandlerService(logger)

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
	adventurerController := controllers.NewAdventurerController(service, logger)

	path := constants.GetAdventurerEndpoint

	// Añade endpoint para retornar un único aventurero por Id
	configureEndpoint(
		router,
		enums.Verb_Get,
		path+"/:"+constants.AdventurerIdParamName,
		errorHandler.HandleError(adventurerController.GetAdventurer))

	// Añade endpoint para eliminar un aventurero
	configureEndpoint(
		router,
		enums.Verb_Delete,
		path+"/:"+constants.AdventurerIdParamName,
		errorHandler.HandleError(adventurerController.DeleteAdventurer))

	// Añade endpoint para crear un nuevo aventurero
	configureEndpoint(
		router,
		enums.Verb_Post,
		path,
		errorHandler.HandleError(adventurerController.CreateAdventurer))

	configureEndpoint(
		router,
		enums.Verb_Post,
		path+"/",
		errorHandler.HandleError(adventurerController.CreateAdventurer))

	// Añade endpoint para actualizar un aventurero
	configureEndpoint(
		router,
		enums.Verb_Update,
		path,
		errorHandler.HandleError(adventurerController.UpdateAdventurer))

	configureEndpoint(
		router,
		enums.Verb_Update,
		path+"/",
		errorHandler.HandleError(adventurerController.UpdateAdventurer))
}

func configureEndpoint(router *gin.Engine, verb enums.Verb, path string, handler gin.HandlerFunc) {
	switch verb {
	case enums.Verb_Get:
		router.GET(path, handler)
	case enums.Verb_Post:
		router.POST(path, handler)
	case enums.Verb_Update:
		router.PUT(path, handler)
	case enums.Verb_Delete:
		router.DELETE(path, handler)
	}
}

// Configura lectura de archivo de configuración con libreria Viper
func configViper() {
	viper.SetConfigName(constants.ConfigFileName)
	viper.SetConfigType(enums.Json)
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
