package services

import (
	"github.com/flickyiyo/pokemon-api/constants"
	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
)

type movesService struct {
	pokemonRepository pokeapi.PokemonRepository
	movesRepository   pokeapi.MoveRepository
}

func NewmovesService(pokemonRepository pokeapi.PokemonRepository, movesRepository pokeapi.MoveRepository) *movesService {
	return &movesService{pokemonRepository, movesRepository}
}

func (self *movesService) FindCommonmoves(request *models.CommonmovesRequest) (*models.movesResponse, error) {
	if len(request.Pokemons) < 2 {
		return nil, constants.NotEnoughPokemons
	}
	var pokemons []*models.Pokemon
	for _, p := range request.Pokemons {
		pokemon, err := self.pokemonRepository.FindPokemon(&p)
		if err != nil {
			return nil, err
		}
		if pokemon == nil {
			return nil, constants.PokemonNotFound(pokemon.Name, pokemon.ID)
		}
		pokemons = append(pokemons, pokemon)
	}

	var commonmoves []models.Move

	return nil, nil
}
