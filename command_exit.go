package main

import (
    "fmt"
    "os"
)

var configExit Config = Config {
    NextUrl: "",
    PrevUrl: nil,
}
func commandExit(config Config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}
