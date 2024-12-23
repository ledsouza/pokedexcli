package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	var locationResp LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResp)
	if err != nil {
		return LocationResponse{}, err
	}

	return locationResp, nil
}
