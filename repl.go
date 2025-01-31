package main

import (
    "strings"
    "fmt"
    "os"
    "bufio"
)

func startPokedex() {
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
        } else {
            err := command.callback(command.config)
            if err != nil {
                fmt.Println("ERROR: cannot perform command")
                fmt.Printf("DETAILS: %v\n", err)
            }
            continue
        }
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
    callback func(config Config) error
    config Config
}

type Config struct {
    NextUrl string
    PrevUrl *string
}

func initializeCommands() map[string]cliCommand {
    return  map[string]cliCommand{
            "exit": {
                name: "exit", 
                description: "Exit the pokedex.", 
                callback: commandExit,
                config: configExit,
            },
            "help": {
                name: "help",
                description: "Displays a help message.",
                callback: commandHelp,
                config: configHelp,
            },
            "map": {
                name: "map", 
                description: `Displays 20 location areas. Each subsequent call 
     displays the next 20 locations.`,
                callback: commandMap,
                config: configMap,
            },
            "mapb": {
                name: "mapb",
                description: `Displays the location areas from the previous 
      page. If user is on the first page, displays that to the user.`,
                callback: commandMapb,
                config: configMap,
            },
        }
}

