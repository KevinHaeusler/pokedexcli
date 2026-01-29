package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/KevinHaeusler/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	var PokegoClient Client = Client{
		httpClient: &http.Client{},
		baseURL:    "https://pokeapi.co/api/v2/",
		cache:      pokecache.NewCache(interval),
	}
	return PokegoClient
}

func (c *Client) Get(path string) ([]byte, error) {
	fullPath := c.baseURL + path
	bytes, ok := c.cache.Get(fullPath)
	if ok {
		return bytes, nil
	}

	resp, err := c.httpClient.Get(fullPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(fullPath, body)
	return body, nil
}
