package pokeapi

import "github.com/flickyiyo/pokemon-api/models"

type PokemonRepository interface {
	FindPokemon(*models.Pokemon) (*models.Pokemon, error)
}

type AbilityRepository interface {
	FindAbility(*models.Ability) (*models.Ability, error)
}
