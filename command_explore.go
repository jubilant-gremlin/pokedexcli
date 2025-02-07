package main

import (
    "fmt"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *Config) error { 
    fullURL := pokeapi.BaseURL + "/location-area/" + cfg.AreaName 
    pokes, err := cfg.pokeapiClient.ListMons(fullURL)
    if err != nil {
        return err
    }

    for _, encounter := range pokes.PokemonEncounters {
        fmt.Println(encounter.Pokemon.Name)
    }

    return nil
}
