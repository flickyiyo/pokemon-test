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

func NewMatchService(matchRepository pokeapi.PokemonRepository, typeRepository pokeapi.TypeRepository) pokeapi.MatchService {
	return &matchService{matchRepository, typeRepository}
}

func (self *matchService) MatchPokemons(request *models.PokemonMatch) (*models.PokemonMatchResponse, error) {

	pokemons := request.Pokemons

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

	typesOfA, err := self.typeRepository.FindTypeList(pokemonA)

	if err != nil {
		return nil, err
	}
	typesOfB, err := self.typeRepository.FindTypeList(pokemonB)
	if err != nil {
		return nil, err
	}

	reasons := hasAdvantage(typesOfA, typesOfB)

	var reasonList []string
	if reasons.DoesDoubleDamage {
		reasonList = append(reasonList, constants.DoesDoubleDamage)
	}
	if reasons.ReceivesDoubleDamage {
		reasonList = append(reasonList, constants.ReceivesDoubleDamage)
	}
	if reasons.DoesHalfDamage {
		reasonList = append(reasonList, constants.DoesHalfDamage)
	}
	if reasons.ReceivesHalfDamage {
		reasonList = append(reasonList, constants.ReceivesHalfDamage)
	}
	if reasons.DoesNoDamage {
		reasonList = append(reasonList, constants.DoesNoDamage)
	}
	if reasons.ReceivesNoDamage {
		reasonList = append(reasonList, constants.ReceivesNoDamage)
	}
	return &models.PokemonMatchResponse{
		Reasons: reasonList,
		Pokemon: *pokemonA,
	}, nil
}

type reasonsStruct struct {
	DoesDoubleDamage     bool
	ReceivesDoubleDamage bool
	DoesHalfDamage       bool
	ReceivesHalfDamage   bool
	DoesNoDamage         bool
	ReceivesNoDamage     bool
}

func hasAdvantage(typesA, typesB []models.Type) *reasonsStruct {
	reasons := reasonsStruct{
		DoesDoubleDamage: false, DoesHalfDamage: false, DoesNoDamage: false,
		ReceivesDoubleDamage: false, ReceivesHalfDamage: false, ReceivesNoDamage: false,
	}
	for _, t := range typesA {
		// Check if types from A are strong against B

		for _, doubleDamage := range t.DamageRelations.DoubleDamageTo {
			for _, tB := range typesB {
				if tB.Name == doubleDamage.Name {
					reasons.DoesDoubleDamage = true
				}
			}
		}
		for _, halfDamage := range t.DamageRelations.HalfDamageFrom {
			for _, tB := range typesB {
				if tB.Name == halfDamage.Name {
					reasons.ReceivesHalfDamage = true
				}
			}
		}
		for _, noDamage := range t.DamageRelations.NoDamageFrom {
			for _, tB := range typesB {
				if tB.Name == noDamage.Name {
					reasons.ReceivesNoDamage = true
				}
			}
		}

		// Check if A is weak against B
		for _, doubleDamage := range t.DamageRelations.DoubleDamageFrom {
			for _, tB := range typesB {
				if tB.Name == doubleDamage.Name {
					reasons.ReceivesDoubleDamage = true
				}
			}
		}
		for _, halfDamage := range t.DamageRelations.HalfDamageTo {
			for _, tB := range typesB {
				if tB.Name == halfDamage.Name {
					reasons.DoesHalfDamage = true
				}
			}
		}
		for _, noDamage := range t.DamageRelations.NoDamageTo {
			for _, tB := range typesB {
				if tB.Name == noDamage.Name {
					reasons.DoesNoDamage = true
				}
			}
		}
	}
	return &reasons
}

func checkIfTypeInlist(typeName string, types []models.Type) bool {
	for _, t := range types {
		if t.Name == typeName {
			return true
		}
	}
	return false
}
