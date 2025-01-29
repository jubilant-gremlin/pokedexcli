package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
        text := scanner.Text()
        text = strings.ToLower(text)
        list := strings.Fields(text)
        cmd := list[0]
        if cmd == "" {
            fmt.Println("error: input is empty")
            return
        }
        if cmd == "exit" {
            commandExit()
        }
    }
}

