package pokeapi

import (
	"github.com/flickyiyo/pokemon-api/models"
)

type MatchService interface {
	MatchPokemons([]models.Pokemon) (*models.PokemonMatchResponse, error)
}

type CommonMovesService interface {
	FindCommonMoves(*models.CommonMovesRequest) (*models.MovesResponse, error)
}

type TypeService interface {
	FindTypeList(*models.PokemonType) ([]models.Type, error)
}
