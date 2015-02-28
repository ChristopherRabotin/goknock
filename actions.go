package main

import "fmt"

// Action is a polymorphic interface. Sequences and their actions must be
// defined in main.go. New actions can be created in actions/.
type Action interface {
    launch()
}

type Echo struct{}

func (act Echo) launch() {
    fmt.Println("Action Echo launched.")
}