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

func NewMovesService(pokemonRepository pokeapi.PokemonRepository, movesRepository pokeapi.MoveRepository) pokeapi.CommonMovesService {
	return &movesService{pokemonRepository, movesRepository}
}

func (self *movesService) FindCommonMoves(request *models.CommonMovesRequest) (*models.MovesResponse, error) {
	if len(request.Pokemons) < 2 {
		return nil, constants.NotEnoughPokemons
	}

	// movesMap := make(map[string]map[string]bool)
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

	return nil, nil
}
