package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := c.cache.Get(url)
	if ok {
		var locationResp LocationResponse
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return LocationResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	var locationResp LocationResponse
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResponse{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
