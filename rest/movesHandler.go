package rest

import (
	"encoding/json"
	"net/http"

	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
)

type MovesHandler interface {
	Post(http.ResponseWriter, *http.Request)
}

type movesHandler struct {
	service pokeapi.CommonMovesService
}

func NewMovesHandler(service pokeapi.CommonMovesService) MovesHandler {
	return &movesHandler{service}
}

func (h *movesHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commonMovesRequest models.CommonMovesRequest
	err := json.NewDecoder(r.Body).Decode(&commonMovesRequest)
	if err != nil {
		w.Write([]byte("Failed to read request body"))
		w.WriteHeader(500)
		return
	}
	responseObj, err := h.service.FindCommonMoves(&commonMovesRequest)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.NewEncoder(w).Encode(responseObj); err != nil {
		w.WriteHeader(500)
	}
}
