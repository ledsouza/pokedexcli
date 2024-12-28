package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, ok := c.cache.Get(url)
	if ok {
		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
