package models

type PokemonMatch struct {
	Pokemon []Pokemon `json:"pokemon"`
}

type PokemonMatchResponse struct {
	Pokemon []Pokemon `json:"pokemon"`
	Reasons []string  `json:"reasons"`
}
