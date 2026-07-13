package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// RespDetailedLocations -
type RespDetailedLocations struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonObject `json:"pokemon"`
}

type PokemonObject struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		Score int `json:"base_stat"`
		Stat  struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
