package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/flickyiyo/pokemon-api/constants"
	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
	"github.com/go-resty/resty/v2"
)

type pokemonRepository struct {
	Client resty.Client
}

func NewPokemonRepository(baseUrl string) pokeapi.PokemonRepository {
	client := resty.New().SetHostURL(baseUrl)
	return &pokemonRepository{Client: *client}
}

func (self *pokemonRepository) FindPokemon(data *models.Pokemon) (*models.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/"
	if &data.ID != nil {
		url = url + fmt.Sprint(data.ID)
	}
	if &data.Name != nil {
		url = url + data.Name
	}
	resp, err := self.Client.R().Get(url)
	if resp.StatusCode() == constants.NotFound {
		return nil, constants.PokemonNotFound(data.Name, data.ID)
	}
	if err != nil {
		return nil, err
	}

	foundPokemon := &models.Pokemon{}

	err = json.Unmarshal(resp.Body(), foundPokemon)

	if err != nil {
		return nil, err
	}

	return foundPokemon, nil
}
