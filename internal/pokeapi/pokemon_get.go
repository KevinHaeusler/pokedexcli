package pokeapi

import "encoding/json"

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	data, err := c.Get("pokemon/" + pokemonName)
	if err != nil {
		return Pokemon{}, err
	}

	var resp Pokemon
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return Pokemon{}, err
	}
	return resp, nil
}
