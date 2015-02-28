package goknock

// Action is a polymorphic interface. Sequences and their actions must be
// defined in main.go. New actions can be created in actions/.
type Action interface {
    launch()
}

type Sequence struct {
    Name string
    ports []int
    action Action
}