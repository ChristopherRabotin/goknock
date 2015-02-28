package main

import "fmt"
import "net"

type Sequence struct {
    name string
    host string
    ports []int
    action Action
    portit int
    conn net.Conn
    ch chan bool
}

// Easy constructor for Sequence structs.
func NewSequence(name string, host string, ports []int, action Action) Sequence{
    return Sequence{name, host, ports, action, -1, nil, make(chan bool)}
}

// Returns the next available port according to the sequence.
func (seq Sequence) nextport() int {
    seq.portit = (seq.portit + 1) % len(seq.ports)
    return seq.ports[seq.portit]
}

// Returns the formatted host and port.
func (seq Sequence) addr() string {
    return fmt.Sprintf("%s:%d", seq.host, seq.nextport())
}

// Accepts the initial knock. Does not have any timeout.
func (seq Sequence) AcceptKnocks() {
    tcpAddr, err := net.ResolveTCPAddr("tcp", seq.addr())
    if err != nil {
        // If we can't resolve the host, we'll panic!
        panic(err)
    }
    netListen, err := net.Listen(tcpAddr.Network(), tcpAddr.String())
    if err != nil {
        // If this port is taken, let's just go to the next one.
        fmt.Errorf("", err)
    }
    conn, err := netListen.Accept()
    if err != nil {
        fmt.Errorf("Client error", err)
    }else{
        fmt.Println("Successful initial connection.")
        // We've got a good connection, let's close it and move to the next one.
        conn.Close()
    }
}

func (seq Sequence) acceptNext(){
    
}