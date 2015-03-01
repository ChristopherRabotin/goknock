package main

import "fmt"
import "net"

// Action is a polymorphic interface. Sequences and their actions must be
// defined in main.go.
type Action interface {
	launch(conn net.Conn)
}

type Echo struct{}

func (act Echo) launch(conn net.Conn) {
	fmt.Println("Action Echo launched.")
}
