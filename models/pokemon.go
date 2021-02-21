package models

type Pokemon struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Types PokemonType `json:"types"`
	Moves []Move      `json:"moves"`
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
