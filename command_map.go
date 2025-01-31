package main

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
    "errors"
)

var configMap Config = Config{
    NextUrl: "",
    PrevUrl: nil,
}

func commandMap(config Config) error {
    if config.NextUrl == "" {
        config.NextUrl = "https://pokeapi.co/api/v2/location-area/"
    }
    res, err := http.Get(config.NextUrl)
    if err != nil {
        fmt.Println("ERROR: could not get response")
        return err
    }
    defer res.Body.Close()
    

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("ERROR: could not read response")
        return err
    }

    list := resourceList{}
    err = json.Unmarshal(body, &list)
    if err != nil {
        fmt.Println("ERROR: could not unmarshal response body")
        return err
    }

    for _, loc := range list.Results {
        fmt.Println(loc.Name)
    }

    configMap.NextUrl = list.Next
    configMap.PrevUrl = list.Previous

    return nil
}

func commandMapb(config Config) error {
    if config.PrevUrl == nil {
        err := errors.New("You're on the first page")
        fmt.Println(err)
        return err
    }
    res, err := http.Get(*config.PrevUrl)
    if err != nil {
        fmt.Println("ERROR: could not get a response")
        return err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("ERROR: could not read response")
        return err
    }

    list := resourceList{}
    err = json.Unmarshal(body, &list)
    if err != nil {
        fmt.Println("ERROR: could not unmarshal response body")
        return err
    }

    for _, loc := range list.Results {
        fmt.Println(loc.Name)
    }

    configMap.PrevUrl = list.Previous
    configMap.NextUrl = list.Next
    return nil
}
