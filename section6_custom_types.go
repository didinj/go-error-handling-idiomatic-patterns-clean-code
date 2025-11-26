package main

import (
    "errors"
    "fmt"
)

type NotFoundError struct {
    Resource string
    ID       int
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
}

func find(id int) error {
    if id == 0 {
        return &NotFoundError{"User", id}
    }
    return nil
}

func main() {
    err := find(0)
    var nf *NotFoundError
    if errors.As(err, &nf) {
        fmt.Println("missing:", nf.Resource, nf.ID)
    }
}
