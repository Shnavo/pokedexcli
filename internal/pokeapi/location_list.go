package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, exists := c.cache.Get(url)

	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespLocations{}, err
		}
		c.cache.Add(url, body)
	}

	locationsResp := RespLocations{}
	err := json.Unmarshal(body, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	return locationsResp, nil
}
