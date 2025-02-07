package main

import (
    "strings"
    "fmt"
    "os"
    "bufio"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokeapi"
    "github.com/jubilant-gremlin/pokedexcli/internal/pokedex"
)

func startPokedex(cfg *Config) {
    // create input scanner and start infinite loop of reading input.
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        // clean input string to slice of lowercase words, cmd = 1st word
        list := CleanInput(scanner.Text())
        if len(list) == 0 {
            continue
        }
        
        cmd := list[0]
        
        // if command is in registry and can be performed without error,
        // perform the command's callback function
        command, ok := initializeCommands()[cmd]
        if !ok {
            fmt.Println("ERROR: Unknown command")
            continue
        }

        // set area name if needed
        if command.name == "explore" {
            if len(list) <= 1 {
                fmt.Println("ERROR: must specify location name to use explore command")
                continue
            }
            cfg.AreaName = list[1]
        }

        // set poke name if needed
        if command.name == "catch" || command.name == "inspect" {
            if len(list) <= 1 {
                fmt.Println("ERROR: must specify pokemon")
                continue
            }
            cfg.PokeName = list[1]
        }

        err := command.callback(cfg)
        if err != nil {
            fmt.Println("ERROR: cannot perform command")
            fmt.Printf("DETAILS: %v\n", err)
        }
        continue
        }
    }


func CleanInput(text string) []string {
    lower := strings.ToLower(text)
    output := strings.Fields(lower)
    return output
}

type cliCommand struct {
    name string
    description string
    callback func(*Config) error
}

type Config struct {
    NextUrl string
    PrevUrl *string
    pokeapiClient pokeapi.Client
    AreaName string
    PokeName string
    dex pokedex.Pokedex
}

func initializeCommands() map[string]cliCommand {
    return  map[string]cliCommand{
            "exit": {
                name: "exit", 
                description: "Exit the pokedex.", 
                callback: commandExit,
            },
            "help": {
                name: "help",
                description: "Displays a help message.",
                callback: commandHelp,
            },
            "map": {
                name: "map", 
                description: `Displays 20 location areas. Each subsequent call 
     displays the next 20 locations.`,
                callback: commandMap,
            },
            "mapb": {
                name: "mapb",
                description: `Displays the location areas from the previous 
      page. If user is on the first page, displays that to the user.`,
                callback: commandMapb,
            },
            "explore": {
                name: "explore",
                description: `Given "explore <area name>" as a command, prints
        a list of pokemon located in that area.`,
                callback: commandExplore,
            },
            "catch": {
                name: "catch",
                description: `Given "catch <pokemon name> as a command, attempts
        to catch the pokemon. Chance of catching the pokemon is determined by pokemon
        base experience. The higher the base experience, the harder the pokemon is to catch`,
                callback: commandCatch,
            },
            "inspect": {
                name: "inspect",
                description: `Given "inspect <pokemon name> as a command, displays
        details about a pokemon if you have caught it before`,
                callback: commandInspect,
            },
            "pokedex": {
                name: "pokedex",
                description: "Displays a list of pokemon the user has caught this session",
                callback: commandPokedex,
            },

        }
}

