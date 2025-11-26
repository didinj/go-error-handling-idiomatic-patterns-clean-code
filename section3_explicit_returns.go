package main

import (
    "fmt"
    "log"
    "os"
)

func readConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("read config: %w", err)
    }
    return data, nil
}

func main() {
    _, err := readConfig("config.json")
    if err != nil {
        log.Println("failed:", err)
    }
}
