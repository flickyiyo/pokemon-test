package models

type Ability struct {
	ID    int           `json:"id"`
	Name  string        `json:"name"`
	Names []AbilityName `json:"names"`
}

type AbilitiesResponse struct {
	Page      int       `json:"page"`
	NumRows   int       `json:"num_rows"`
	Abilities []Ability `json:"abilities"`
}

type AbilityName struct {
	Name     string   `json:"name"`
	Language language `json:"language"`
}

type language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
