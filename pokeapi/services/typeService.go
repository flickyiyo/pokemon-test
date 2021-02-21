package services

import (
	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
)

type typeService struct {
	typeRepository pokeapi.TypeRepository
}

func NewTypeService(typeRepository *pokeapi.TypeRepository) *typeService {
	return &typeService{*typeRepository}
}

func (self *typeService) FindTypes(pokemon *models.Pokemon) (*models.Type, error) {
	return nil, nil
}
