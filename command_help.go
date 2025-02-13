package main

import (
    "fmt"
)

func commandHelp(cfg *Config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("")

    // generate up-to-date command registry and print command descriptions
    registry := initializeCommands()
    for _, cmd := range registry {
        fmt.Printf("%v: %v\n", cmd.name, cmd.description)
    }
    return nil
} 
