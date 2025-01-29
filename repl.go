package main

import (
    "strings"
)

func CleanInput(text string) []string {
    lower := strings.ToLower(text)
    output := strings.Fields(lower)
    return output
}
