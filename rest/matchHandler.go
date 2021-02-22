package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
	"github.com/go-chi/chi"
)

type MatchHandler interface {
	Get(http.ResponseWriter, *http.Request)
	// Post(http.ResponseWriter, *http.Request)
}

type matchHandler struct {
	matchService pokeapi.MatchService
	typeService  pokeapi.TypeService
}

func NewMatchHandler(matchService pokeapi.MatchService, typeService pokeapi.TypeService) MatchHandler {
	return &matchHandler{matchService, typeService}
}

func (h *matchHandler) Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	idA := chi.URLParam(r, "pokemonA")
	idB := chi.URLParam(r, "pokemonB")

	var pokemonA models.Pokemon
	var pokemonB models.Pokemon

	if val, err := strconv.Atoi(idA); err != nil {
		pokemonA.ID = val
	} else {
		pokemonA.Name = idA
	}

	if val, err := strconv.Atoi(idB); err != nil {
		pokemonB.ID = val
	} else {
		pokemonB.Name = idA
	}

	response, err := h.matchService.MatchPokemons(&models.PokemonMatch{
		Pokemons: []models.Pokemon{pokemonA, pokemonB},
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Fatal("response object not get")
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
	}

}
