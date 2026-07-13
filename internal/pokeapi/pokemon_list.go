package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(area string) (RespDetailedLocations, error) {
	url := baseURL + "/location-area" + "/" + area
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespDetailedLocations{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespDetailedLocations{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDetailedLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDetailedLocations{}, err
	}

	pokemonResp := RespDetailedLocations{}
	if err = json.Unmarshal(dat, &pokemonResp); err != nil {
		return RespDetailedLocations{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
