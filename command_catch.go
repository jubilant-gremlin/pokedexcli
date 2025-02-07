package main

import (
    "fmt"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
    "math/rand"
)

func commandCatch(cfg *Config) error {
    fmt.Printf("Throwing a Pokeball at %v...\n", cfg.PokeName)
    fullURL := pokeapi.BaseURL + "/pokemon/" + cfg.PokeName
    poke, err := cfg.pokeapiClient.MonInfo(fullURL)
    if err != nil {
        return err
    }

    // catch rate logic
    n := rand.Intn(255)
    f := (poke.BaseExperience * 255 * 4) / (poke.BaseExperience * 7)
    if f >= n {
        fmt.Printf("%v was caught!\n", poke.Name)
        fmt.Printf("%v can now be inspected with the inspect command\n", poke.Name)
        cfg.dex.PokedexEntries[cfg.PokeName] = poke
    } else {
        fmt.Println(poke.BaseExperience, f, n)
        fmt.Printf("%v escaped!\n", poke.Name)
    }

    return nil
}
