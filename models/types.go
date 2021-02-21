package models

type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
