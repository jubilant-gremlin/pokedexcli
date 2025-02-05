package main

import (
    "errors"
    "fmt"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {     
    fullURL := pokeapi.BaseURL + "/location-area"
    if cfg.NextUrl == "" {
        cfg.NextUrl = fullURL
    }
    locs, err := cfg.pokeapiClient.ListLocations(&cfg.NextUrl)
    if err != nil {
        return err
    }

    cfg.NextUrl = locs.Next
    cfg.PrevUrl = locs.Previous

    for _, loc := range locs.Results {
        fmt.Println(loc.Name)
    }
    return nil
}

func commandMapb(cfg *Config) error {
    if cfg.PrevUrl == nil {
        return errors.New("you're on the first page")
    }

    locs, err := cfg.pokeapiClient.ListLocations(cfg.PrevUrl)
    if err != nil {
        return err
    }

    cfg.NextUrl = locs.Next
    cfg.PrevUrl = locs.Previous

    for _, loc := range locs.Results {
        fmt.Println(loc.Name)
    }
    return nil
}
