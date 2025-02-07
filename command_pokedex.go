package main

import (
    "fmt"
)

func commandPokedex(cfg *Config) error {
    fmt.Println("YOUR POKEDEX:")
    if len(cfg.dex.PokedexEntries) == 0 {
        fmt.Println("POKEDEX IS EMPTY")
    }
    for _, poke := range cfg.dex.PokedexEntries {
        fmt.Printf("    -%v\n", poke.Name)
    }
    return nil
}
