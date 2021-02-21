package constants

import (
	"errors"
)

var (
	PokemonNotFound = func(pokemonName string, pokemonID int) error {
		return errors.New("Pokemon not found.\n" +
			"Pokemon Name: " + pokemonName + "\n Pokemon ID: " + string(pokemonID))
	}
	NotEnoughPokemons       = errors.New("Not enough Pokemons to perform request")
	PokemonQuantityExceeded = errors.New("List of pokemons bigger than expected")
)
