package services

import (
	"github.com/Brackistar/guild-hall-backend/models"
)

type MockAdventurerService struct {
	totalAdventurerCounter uint64
}

func NewMockAdventurerService() *MockAdventurerService {
	return &MockAdventurerService{
		totalAdventurerCounter: 1,
	}
}

func (s MockAdventurerService) GetAdventurer(id *uint64) (models.Adventurer, error) {
	result := models.Adventurer{
		Id:   1,
		Name: "Test adventurer",
		Rank: models.AdvRank{
			Id:          1,
			Name:        "Test-Rank",
			Order:       1,
			PointsToGet: 0,
		},
		Class: models.AdvClass{
			Id:   1,
			Name: "Test-Class",
		},
	}

	return result, nil
}

func (s MockAdventurerService) CreateAdventurer(adventurer *models.Adventurer) (uint64, error) {
	s.totalAdventurerCounter++
	return s.totalAdventurerCounter, nil
}

func (s MockAdventurerService) DeleteAdventurer(id *uint64) error {
	return nil
}

func (s MockAdventurerService) UpdateAdventurer(adventurer *models.Adventurer) (models.Adventurer, error) {
	return *adventurer, nil
}
