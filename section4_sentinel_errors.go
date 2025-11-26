package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) (string, error) {
    if id == 0 {
        return "", ErrNotFound
    }
    return "User", nil
}

func main() {
    _, err := findUser(0)
    if errors.Is(err, ErrNotFound) {
        fmt.Println("user not found")
    }
}
