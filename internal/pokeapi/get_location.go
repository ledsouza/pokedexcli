package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (ExploreLocationResponse, error) {
	url := baseURL + "/location-area/" + location

	data, ok := c.cache.Get(url)
	if ok {
		exploreLocationResp := ExploreLocationResponse{}
		err := json.Unmarshal(data, &exploreLocationResp)
		if err != nil {
			return ExploreLocationResponse{}, err
		}
		return exploreLocationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreLocationResponse{}, err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return ExploreLocationResponse{}, err
	}

	exploreLocationResp := ExploreLocationResponse{}
	err = json.Unmarshal(data, &exploreLocationResp)
	if err != nil {
		return ExploreLocationResponse{}, err
	}

	c.cache.Add(url, data)

	return exploreLocationResp, nil
}
