package pokeapi

import (
	"github.com/flickyiyo/pokemon-api/models"
)

type MatchService interface {
	MatchPokemons([]models.Pokemon) (*models.Pokemon, error)
}

type CommonMovesService interface {
	FindCommonMoves(*models.CommonMovesRequest) (*models.MovesResponse, error)
}
