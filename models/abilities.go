package models

type Move struct {
	ID    int        `json:"id"`
	Name  string     `json:"name"`
	Names []MoveName `json:"names"`
}

type MovesResponse struct {
	Page    int    `json:"page"`
	NumRows int    `json:"num_rows"`
	Moves   []Move `json:"moves"`
}

type CommonMovesRequest struct {
	LangID   int       `json:"lang_id"`
	LangName string    `json:"lang_name"`
	NumRows  int       `json:"num_rows"`
	Pokemons []Pokemon `json:"pokemons"`
}

type MoveName struct {
	Name     string   `json:"name"`
	Language language `json:"language"`
}

type language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
