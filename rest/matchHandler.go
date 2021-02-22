package rest

import (
	"net/http"

	"github.com/flickyiyo/pokemon-api/pokeapi"
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

}
