package pokeapi

import (
	"github.com/flickyiyo/pokemon-api/models"
)

type MatchService interface {
	MatchPokemons([]models.Pokemon) (*models.Pokemon, error)
	FindAbilities([]models.Pokemon) (*models.AbilitiesResponse, error)
}
