package models

type Type struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	DoubleDamageTo   []typeOverview `json:"double_damage_to"`
	HalfDamageTo     []typeOverview `json:"half_damage_to"`
	DoubleDamageFrom []typeOverview `json:"double_damage_from"`
	HalfDamageFrom   []typeOverview `json:"half_damage_from"`
	NoDamageTo       []typeOverview `json:"no_damage_to"`
	NoDamageFrom     []typeOverview `json:"no_damange_from"`
}

type PokemonType struct {
	ID   int          `json:"id"`
	Slot int          `json:"slot"`
	Type typeOverview `json:"type"`
}

type typeOverview struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
