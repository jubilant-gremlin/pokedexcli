package main

import (
    "fmt"
    "errors"
)

func commandInspect(cfg *Config) error {
    dexEntry, ok := cfg.dex.PokedexEntries[cfg.PokeName]
    if !ok{
        err := errors.New("Dex entry not found, must catch pokemon to inspect")
        return err
    }
    fmt.Printf("NAME: %v\n", dexEntry.Name)
    fmt.Printf("HEIGHT: %v\n", dexEntry.Height)
    fmt.Printf("WEIGHT: %v\n", dexEntry.Weight)
    fmt.Println("STATS:")
    for _, stat := range dexEntry.Stats {
        fmt.Printf("    -%v: %v\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("TYPES:")
    for _, t := range dexEntry.Types {
        fmt.Printf("    -%v\n", t.Type.Name)
    }
    return nil
}

