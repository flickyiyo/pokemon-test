package constants

import (
	"errors"
)

var (
	PokemonNotFound = func(pokemon string) error {
		return errors.New("Pokemon not found" + pokemon)
	}
)
