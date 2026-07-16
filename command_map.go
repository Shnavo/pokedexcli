package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
