package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
