package pokeapi

import "encoding/json"

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
