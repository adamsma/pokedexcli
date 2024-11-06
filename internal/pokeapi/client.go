package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/adamsma/pokedexcli/internal/pokecache"
)

const apiurl = "https://pokeapi.co/api/v2/"

func do(endpoint string, obj interface{}, c *pokecache.Cache) error {

	val, found := c.Get(endpoint)

	if !found {

		res, err := http.Get(apiurl + endpoint)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}

		go func() {
			c.Add(endpoint, data)
		}()
		val = data

	}

	return json.Unmarshal(val, &obj)

}
