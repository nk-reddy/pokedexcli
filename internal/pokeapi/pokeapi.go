package pokeapi

import (
	"time"

	"github.com/nk-reddy/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

var cache = pokecache.NewCache(5 * time.Second)
