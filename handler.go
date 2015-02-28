package main

import "net"
import "fmt"

func HandleConnection(conn net.Conn){
    // seqchan receives true if the next knock is received.
    // It is reset at every successful knock.
    seqchan := make(chan bool, 1)
    // We've got a connection! Let's close it and go the next port.
    fmt.Print(conn)
}