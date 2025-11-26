package main

import (
    "fmt"
    "golang.org/x/sync/errgroup"
)

func taskA() error { return fmt.Errorf("taskA failed") }
func taskB() error { return nil }

func main() {
    g := new(errgroup.Group)

    g.Go(func() error { return taskA() })
    g.Go(func() error { return taskB() })

    if err := g.Wait(); err != nil {
        fmt.Println("error:", err)
    }
}
