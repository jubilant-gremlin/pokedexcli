package main

import (
    "time"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokedex"

)
func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
    dex := pokedex.NewDex()
    cfg := &Config{
        pokeapiClient: pokeClient,
        dex: dex,
    }
    startPokedex(cfg)
}

