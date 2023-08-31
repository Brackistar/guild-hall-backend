package controllers

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/Brackistar/guild-hall-backend/constants"
	"github.com/Brackistar/guild-hall-backend/interfaces"
	"github.com/Brackistar/guild-hall-backend/models"
	"github.com/gin-gonic/gin"
)

type AdventurerController struct {
	adventurerService interfaces.IAdventurerService
	logger            *logrus.Logger
}

func NewAdventurerController(service interfaces.IAdventurerService, logger *logrus.Logger) *AdventurerController {
	return &AdventurerController{
		adventurerService: service,
		logger:            logger,
	}
}

func (controller *AdventurerController) GetAdventurer(c *gin.Context) (string, error) {
	id, err := strconv.ParseUint(c.Param(constants.AdventurerIdParamName), 10, 64)

	if err != nil {
		return badRequestResponse(err)
	}

	adventurer, err := controller.adventurerService.GetAdventurer(&id)

	if err != nil {
		return internalServerResponse(err)
	}

	c.JSON(http.StatusOK, adventurer)

	return "", nil
}

func (controller *AdventurerController) CreateAdventurer(c *gin.Context) (string, error) {
	var adventurer models.Adventurer

	if err := c.BindJSON(&adventurer); err != nil {
		return badRequestResponse(err)
	}

	newId, err := controller.adventurerService.CreateAdventurer(&adventurer)

	if err != nil {
		return internalServerResponse(err)
	}

	c.JSON(http.StatusOK, newId)

	return "", nil
}

func (controller *AdventurerController) UpdateAdventurer(c *gin.Context) (string, error) {
	var adventurer models.Adventurer

	if err := c.BindJSON(&adventurer); err != nil {
		return badRequestResponse(err)
	}

	result, err := controller.adventurerService.UpdateAdventurer(&adventurer)

	if err != nil {
		return internalServerResponse(err)
	}

	c.JSON(http.StatusOK, result)

	return "", nil
}

func (controller *AdventurerController) DeleteAdventurer(c *gin.Context) (string, error) {
	id, err := strconv.ParseUint(c.Param(constants.AdventurerIdParamName), 10, 64)

	if err != nil {
		return badRequestResponse(err)
	}

	err = controller.adventurerService.DeleteAdventurer(&id)

	if err != nil {
		return internalServerResponse(err)
	}

	c.Status(http.StatusAccepted)

	return "", nil
}

func badRequestResponse(err error) (string, error) {
	return http.StatusText(http.StatusBadRequest), err
}

func internalServerResponse(err error) (string, error) {
	return http.StatusText(http.StatusInternalServerError), err
}
