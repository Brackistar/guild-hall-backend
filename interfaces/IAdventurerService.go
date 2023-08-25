package interfaces

import "github.com/Brackistar/guild-hall-backend/models"

type IAdventurerService interface {
	GetAdventurer(id *uint64) (models.Adventurer, error)
	CreateAdventurer(adventurer *models.Adventurer) (uint64, error)
	DeleteAdventurer(id *uint64) error
	UpdateAdventurer(adventurer *models.Adventurer) (models.Adventurer, error)
}
