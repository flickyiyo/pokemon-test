package pokeapi

import (
	"github.com/flickyiyo/pokemon-api/models"
)

type MatchService interface {
	MatchPokemons(*models.PokemonMatch) (*models.PokemonMatchResponse, error)
}

type CommonMovesService interface {
	FindCommonMoves(*models.CommonMovesRequest) (*models.MovesResponse, error)
}

type TypeService interface {
	FindTypeList(*models.Pokemon) ([]models.Type, error)
}
