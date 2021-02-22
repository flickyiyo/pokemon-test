package models

type PokemonMatch struct {
	Pokemons []Pokemon `json:"pokemon"`
}

// PokemonMatchResponse is the response object for match endpint
type PokemonMatchResponse struct {
	Pokemon Pokemon  `json:"pokemon"`
	Reasons []string `json:"reasons"`
}
