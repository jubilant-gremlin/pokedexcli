package pokedex

import (
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
    PokedexEntries map[string]pokeapi.Pokemon
}

func NewDex() Pokedex{
    new := Pokedex {
        PokedexEntries: make(map[string]pokeapi.Pokemon),
    }
    return new
}
