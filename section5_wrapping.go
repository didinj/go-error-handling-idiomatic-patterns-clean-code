package main

import (
    "errors"
    "fmt"
    "os"
)

func loadConfig() error {
    _, err := os.ReadFile("config.json")
    if err != nil {
        return fmt.Errorf("load config: %w", err)
    }
    return nil
}

func main() {
    err := loadConfig()
    if err != nil {
        fmt.Println(err)
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("file missing")
        }
    }
}
