package models

type Move struct {
	// ID    int        `json:"id"`
	Name  string     `json:"name"`
	Names []MoveName `json:"names"`
}

type MovesResponse struct {
	Page    int       `json:"page"`
	NumRows int       `json:"num_rows"`
	Moves   []MoveDto `json:"moves"`
}

type MoveDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
}

type CommonMovesRequest struct {
	LangName string    `json:"lang_name"`
	NumRows  int       `json:"num_rows"`
	NumPage  int       `json:"num_page"`
	Pokemons []Pokemon `json:"pokemons"`
}

type MoveName struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
