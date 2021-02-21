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

func (self *typeService) FindTypeList(pokemon *models.Pokemon) ([]models.Type, error) {
	typeList, err := self.typeRepository.FindTypeList(pokemon)
	if err != nil {
		return nil, err
	}
	return typeList, nil
}
