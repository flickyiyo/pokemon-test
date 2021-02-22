package models

type Pokemon struct {
	ID    int                   `json:"id"`
	Name  string                `json:"name"`
	Types []PokemonType         `json:"types"`
	Moves []PokemonMovesWrapper `json:"moves"`
}

type PokemonMovesWrapper struct {
	Move PokemonMove `json:"move"`
}

type PokemonRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokemonMove struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
