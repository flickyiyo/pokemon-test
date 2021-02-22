package pokeapi

import "github.com/flickyiyo/pokemon-api/models"

type PokemonRepository interface {
	FindPokemon(*models.Pokemon) (*models.Pokemon, error)
}

type MoveRepository interface {
	FindMove(*models.Move) (*models.Move, error)
}

type TypeRepository interface {
	// FindType(*models.Type) (*models.Type, error)
	FindTypeList(*models.Pokemon) ([]models.Type, error)
}
