package services

import (
	"sort"

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
	var lang string
	if request.LangName == "" || &request.LangName == nil {
		lang = "en"
	} else {
		lang = request.LangName
	}
	movesMap := make(map[string][]string)
	pokemons := request.Pokemons
	for _, p := range pokemons {
		pokemon, err := self.pokemonRepository.FindPokemon(&p)
		if err != nil {
			return nil, err
		}
		if pokemon == nil {
			return nil, constants.PokemonNotFound(pokemon.Name, pokemon.ID)
		}
		for _, move := range pokemon.Moves {
			if movesMap[move.Move.Name] != nil {
				movesMap[move.Move.Name] = append(movesMap[move.Move.Name], pokemon.Name)
			} else {
				movesMap[move.Move.Name] = []string{pokemon.Name}
			}
		}
	}

	var allSharedMoves []string
	for k, v := range movesMap {
		if len(v) == len(pokemons) {
			allSharedMoves = append(allSharedMoves, k)
		}
	}
	sort.Strings(allSharedMoves)

	numPage := request.NumPage
	if request.NumPage == 0 || &request.NumPage == nil {
		numPage = 1
	}
	slicedSharedMoves := allSharedMoves[numPage*10-10 : numPage*10]

	var dtoSharedMoves []models.MoveDto
	for _, sharedMove := range slicedSharedMoves {
		foundMove, err := self.movesRepository.FindMove(&models.Move{Name: sharedMove})
		if err != nil {
			return nil, err
		}
		mv := getMoveOnLang(foundMove, lang)
		dtoSharedMoves = append(dtoSharedMoves, *mv)
	}

	return &models.MovesResponse{
		Page:    request.NumPage,
		Moves:   dtoSharedMoves,
		NumRows: 10,
	}, nil
}

func getMoveOnLang(move *models.Move, lang string) *models.MoveDto {
	for _, moveName := range move.Names {

		if moveName.Language.Name == lang {
			return &models.MoveDto{
				Name:     moveName.Name,
				Language: lang,
			}
		}
	}
	return nil
}
