package models

type Pokemon struct {
	ID   int16       `json:"id"`
	Name string      `json:"name"`
	Type PokemonType `json:"types"`
}

type PokemonType struct {
	ID   int         `json:"id"`
	Slot int         `json:"slot"`
	Type typeDetails `json:"type"`
}

type typeDetails struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
