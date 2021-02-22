package repositories

import (
	"encoding/json"
	"errors"

	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
	"github.com/go-resty/resty/v2"
)

type typeRepository struct {
	Client resty.Client
}

func NewTypeRepository(baseUrl string) pokeapi.TypeRepository {
	client := resty.New().SetHostURL(baseUrl)
	return &typeRepository{Client: *client}
}

func (self *typeRepository) FindTypeList(pokemon *models.Pokemon) ([]models.Type, error) {
	var slotTypes []models.Type

	for _, slot := range pokemon.Types {
		t := findType(slot.Type.Name, self.Client)
		if t == nil {
			return nil, errors.New("Type not found")
		}
		slotTypes = append(slotTypes, *t)
	}

	return slotTypes, nil
}

func findType(typeName string, client resty.Client) *models.Type {
	resp, err := client.R().Get("/type/" + typeName)
	if err != nil {
		return nil
	}
	slotType := &models.Type{}

	err = json.Unmarshal(resp.Body(), slotType)
	if err != nil {
		return nil
	}

	return slotType
}
