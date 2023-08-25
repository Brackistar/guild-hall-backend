package controllers

import (
	"net/http"

	"github.com/sirupsen/logrus"

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

func (controller *AdventurerController) GetAdventurer(c *gin.Context) {
	var id uint64
	if err := c.Bind(&id); err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	adventurer, err := controller.adventurerService.GetAdventurer(&id)

	if err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, adventurer)
}

func (controller *AdventurerController) CreateAdventurer(c *gin.Context) {
	var adventurer models.Adventurer

	if err := c.BindJSON(&adventurer); err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newId, err := controller.adventurerService.CreateAdventurer(&adventurer)

	if err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newId)
}

func (controller *AdventurerController) UpdateAdventurer(c *gin.Context) {
	var adventurer models.Adventurer

	if err := c.BindJSON(&adventurer); err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := controller.adventurerService.UpdateAdventurer(&adventurer)

	if err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *AdventurerController) DeleteAdventurer(c *gin.Context) {
	var id uint64

	if err := c.BindJSON(&id); err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := controller.adventurerService.DeleteAdventurer(&id)

	if err != nil {
		controller.logger.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusAccepted)
}
