package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2/"

func do(endpoint string, obj interface{}) error {

	res, err := http.Get(apiurl + endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		return err
	}

	return decoder.Decode(&obj)

}
