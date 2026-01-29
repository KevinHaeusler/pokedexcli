package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocationAreas(path string) (LocationAreaResponse, error) {
	data, err := c.Get(path)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var resp LocationAreaResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return resp, nil
}

func (c *Client) GetLocationArea(path string) (LocationArea, error) {
	data, err := c.Get(path)
	if err != nil {
		return LocationArea{}, err
	}

	var resp LocationArea
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return LocationArea{}, err
	}

	return resp, nil
}
