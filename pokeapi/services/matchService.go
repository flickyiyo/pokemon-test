package services

import (
	"github.com/flickyiyo/pokemon-api/constants"
	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
)

type matchService struct {
	pokemonRepository pokeapi.PokemonRepository
	typeRepository    pokeapi.TypeRepository
}

func NewMatchService(matchRepository *pokeapi.PokemonRepository, typeRepository *pokeapi.TypeRepository) *matchService {
	return &matchService{*matchRepository, *typeRepository}
}

func (self *matchService) MatchPokemons(pokemons []models.Pokemon) (*models.Pokemon, error) {
	if len(pokemons) < 2 {
		return nil, constants.NotEnoughPokemons
	}
	if len(pokemons) > 2 {
		return nil, constants.PokemonQuantityExceeded
	}
	pokemonA, err := self.pokemonRepository.FindPokemon(&pokemons[1])
	if err != nil {
		return nil, err
	}
	pokemonB, err := self.pokemonRepository.FindPokemon(&pokemons[0])
	if err != nil {
		return nil, err
	}

	return nil, nil
}
