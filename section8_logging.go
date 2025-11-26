package main

import (
    "fmt"
    "log"
    "os"
)

func readFile(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("readFile: %w", err)
    }
    return data, nil
}

func main() {
    _, err := readFile("missing.txt")
    if err != nil {
        log.Println("failed:", err)
    }
}
