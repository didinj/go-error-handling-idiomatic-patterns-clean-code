package main

import (
    "fmt"
)

func risky() {
    panic("boom")
}

func safeRun() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    risky()
    return nil
}

func main() {
    fmt.Println(safeRun())
}
