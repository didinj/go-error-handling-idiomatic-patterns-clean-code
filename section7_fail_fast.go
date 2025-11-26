package main

import (
    "errors"
    "fmt"
)

var ErrInvalidData = errors.New("invalid data")

func process(data int) error {
    if data <= 0 {
        return ErrInvalidData
    }
    return nil
}

func main() {
    if err := process(-1); err != nil {
        fmt.Println(err)
    }
}
